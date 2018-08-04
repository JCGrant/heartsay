package heartsay

import (
	"github.com/JCGrant/heartsay/renderer"
)

var hearts = []string{
	`
          ####   ####
        ##    # #    ##
      ##       #       ##
     #                   #
    #                     #
    #                     #
    #                     #
     #                   #
      #                 #
       #               #
        ##           ##
          #         #
           ##     ##
             ## ##
               #
`,
	`
        #####     #####
      ##     #   #     ##
    ##        # #        ##
  ##           #           ##
 #                           #
#                             #
#                             #
#                             #
#                             #
 #                           #
  #                         #
   #                       #
    ##                   ##
      #                 #
       ##             ##
         #           #
          ##       ##
            ##   ##
              ###
`,
}

var textBounds = renderer.NewRect(7, 6, 23, 8)
