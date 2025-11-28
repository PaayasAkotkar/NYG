©️ copyright 2025
all rights reversed
🌐                                     NYG - BACKEND                           🌐
NOTE ⚔️: docker file testing is not done so if you want to try and test it you can

index:
        1.collection
        2.credentials
        3.game
        4.graphql-chat-app
        5.msg-webhook
        6.patch
        7.post
        8.profile
        9.protos
        10.sql-manager
        11.updates

1. collection
             sse implmented database of whole which is the root of the game
2. credentials
              motivation🔥:  a rest api to update game       
                             progress [credits, power-up ugpgrades]
              design⚙️: simple and robust in reading and some basic knowledge of json been applied
3. game
      motivation🔥:  a websocket server who is the child of whole packages
        design⚙️:
                kept simple and robust 
                the gofiber provides that much as much required to get it done
                the hub is quite clean to look at which is just reading a book

4. graphql-chat-app
         motivation 🔥: name says it all and the reason behind graphql is because i dont want to go through long polling drama or neither i wanted to add one more websocket or i dont want to spend time with sse
         note: currenlty the testing of chat saving has not been implmented so make sure to test it 
         design⚙️: so i am fan of pub-sub system since i research about them
         so i tried implmenting my own
         so its purely based on pub-sub 

5. msg-webhook
        motivation 🔥: i got motivation from github itself like alot uses webhook so yeah i created it incase if you were to send the message to the client 
        design⚙️: its a webhook nothing special to say right now

6. patch
      motivation 🔥: creating a login version of the backend
      design ⚙️: this is the core of the whole
      without this the whole backend will fail
      the design is quite simple it just patches the request 

7. post
    motivation 🔥: at the beginning the post was actully implemted for frontend it has nothing to do with what it does now
    post is grpc service which is the core and connects the game package and someother [yeah i wrote so many codes that now i dont know where the post is used my bad 🥱]
    design⚙️: quite robust and basic mysql stuff that just fetches the data and post's it via graphql

8. profile
           motivation 🔥: since the beginning this folder was actually created for the player's appreance like pic, game theme and blah blah
           as the time progressed this is the core for books and broadcasting the player's progress
           design ⚙️: quite complicated but again its easy
           the design is based on my-sql replication [magical stuff 🪄], my-sql itself for verification and fetching, the main is noother than graphql [modern dev baby] 
           after combining all this a pub-sub system been designed 
           i mean i said its complicated because the pub-sub system is not desgined how graphql-chat-... system been implmented

9. protos
       motivation 🔥: travel amongst backend
                     so at the begining whole was desinged based on sse but its just coroutine issues bunch of shit things happened that you dont want to know
                     so i discovered grpc and its just made things easy
                     in short: i used grpc to not play with golang corutine implment sse
        desgin⚙️: there's nothing to implmented tho what design

10. sql-manager
        motivation 🔥: you dont want to know what acutally happened earlier so that's why this badass package is created
        design⚙️: a micro-service to minimally created the services which are needed to get this project done

11. update
        motivation 🔥: updates the game results of the player such as rating and other stuffs
        design⚙️: quite robust nothing special if you know mysql


Walkthrough 🚶:
                just run the compile.bat file 
                please NOTE it will open all 9 terminals than 8 separate tabs
                
# Whats left:
              x: i am still unsure about the validation of bets and repeated guess token
                 a chat save option

                0. clash mode effect 
                1. login system
                2. security system
                3. kafka for spectating mode
                4. cluster for clash match making [typically planning to use google for games]
                5. database
                6. server
                7. end-to-end testing
                8. load-balancing

# problems:
            post folder will end up having close once the mysql conn limit reached