 {
     "jsonrpc": "2.0",
     "method": "deploy",
     "params": {
         "type": 1,
         "chaincodeID": {
             "path": "https://github.com/eoinwoods/learn-chaincode/start"
         },
         "ctorMsg": {
             "function": "init",
             "args": [
                 "hi there"
             ]
         },
         "secureContext": "user_type1_0"
     },
     "id": 1
 }