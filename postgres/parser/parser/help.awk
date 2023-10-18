# Copyright 2023 Dolthub, Inc.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

# Copyright 2017 The Cockroach Authors.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http:#www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or
# implied. See the License for the specific language governing
# permissions and limitations under the License.
#

# This file extracts inline help specifications from the grammar file
# (sql.y) into a Go map[string]HelpMessageBody suitable for building
# messages towards clients.
#
# See README.md for details.

# Hint, to help read awk code: remember in awk strings are 1-indexed!

BEGIN {
    # inhelp serves to assert below that the special
    # directives follow an initial "%Help:" marker,
    # and do not appear after "%End".
    inhelp = 0
    # helpkey is the current help key, used for error messages.
    helpkey = ""
    # instring indicates whether we are currently accumulating
    # text for a Go string. If this is set, then we
    # first must close the string before generating more code.
    instring = 0

    # Header in the generated code.
    print "// Code generated by help.awk. DO NOT EDIT."
    print "// GENERATED FILE DO NOT EDIT"
    print
    print "package parser"
    print
    print "var helpMessages = map[string]HelpMessageBody{"
}

/^\/\/ %Help:/ && inhelp == 1 { printf "%d: unexpected: %s (last: %s)", NR, $0, helpkey >"/dev/stderr";    exit 1 }

/^\/\/ %Help:/ && inhelp == 0 {
    inhelp = 1
    instring = 0
    # // %Help:   SOMEKEY - sometext
    #          ^^^^^^^^^^^^^^^^^^^^^  take this
    helprest = substr($0, index($0, ":")+1)

    # Remove spaces at the beginning.
    sub("^[ \t]*", "", helprest)

    # Now split the key from the summary, if there's one.
    helpkey = helprest
    summary = ""
    has_summary = index(helprest, "-")
    if (has_summary > 0) {
        # SOMEKEY - sometext
        #          ^^^^^^^^^ take this
        summary = substr(helprest, has_summary+1)
        # Remove spaces at the beginning.
        sub("^[ \t]*", "", summary)
        # SOMEKEY     - sometext
        # ^^^^^^^^^^^ take this
        helpkey = substr(helprest, 1, has_summary-1)
    }
    # Remove trailing spaces.
    sub("[ \t]*$", "", helpkey)

    # Output a line marker so that errors in the generated code
    # properly refer to the original file.
    print "  //line sql.y:", NR

    # Generate the code.
    printf "  `%s`: {\n", helpkey
    if (summary != "") {
        printf "    ShortDescription: `%s`,\n", summary
    }
    next
}

/^\/\/ %/ && inhelp == 0 { printf "%d: unexpected: %s (last: %s)", NR, $0, helpkey >"/dev/stderr"; exit 1 }

/^\/\/ %Category:/ && inhelp == 1 {
    # // %Category:   HELLO WORLD
    #              ^^^^^^^^^^^^^^ take this
    key = substr($0, index($0, ":")+1)
    # Remove heading spaces.
    sub("^[ \t]*", "", key)
    # If we were accumulating a string, close it.
    if (instring == 1) {
        print "`,"
    }
    # Emit the category marker.
    print "    //line sql.y:", NR
    printf "    Category: h%s,\n", key
    instring = 0
    next
}

/^\/\/ %[^:]*:/ && inhelp == 1 {
    # // %SomeKey:   hello world
    #     ^^^^^^^ take this
    key = substr($0, index($0, "%")+1, index($0,":")-index($0, "%")-1)
    # // %SomeKey:   hello world
    #             ^^^^^^^^^^^^^^ take this
    rest = substr($0, index($0, ":")+1)
    # Remove heading spaces.
    sub("^[ \t]*", "", rest)
    # If we were accumulating a string, close it.
    if (instring == 1) {
        print "`,"
    }
    # Emit the key/string pair.
    print "    //line sql.y:", NR
    printf "    %s: `%s\n", key, rest
    instring = 1
    next
}

(/^[^\/]/ || /^\/\/ %End/) && inhelp == 1 {
    # If we were accumulating a string, close it.
    if (instring == 1) {
        print "`,"
    }
    # Finish the Go dictionary entry.
    print "  },"
    inhelp = 0
    instring = 0
    next
}

/^\/\// && inhelp == 1 {
    # We are accumulating content, whatever it is.
    # Put it in the generated code.
    print substr($0, 4)
    next
}

END {
    if (inhelp == 1) {
        printf "missing %% End (last: %s)\n", helpkey >"/dev/stderr"
        exit 1
    }
    print "}"
}
