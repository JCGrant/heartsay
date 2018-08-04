package heart

import (
	"github.com/JCGrant/heart/renderer"
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

var textBounds = renderer.NewRect(7, 4, 23, 6)
