var express = require('express');
var app = express();
//const query = require('./fabcar/query.js')

// you can make a request to a database here and retrieve some data
// but for this example, we are using a static object of user information
var Users = {
    
      'David': {
        age: 52,
        occupation: 'Professor',
        hobby: 'Swimming'
      },
    
      'Robert': {
        age: 34,
        occupation: 'Engineer',
        hobby: 'Running'
      },
    
      'Jane': {
        age: 28,
        occupation: 'Nurse',
        hobby: 'Chess'
      }
    
    };
    
    // route pages
    app.get('/', function (req, res) {
      res.send('Main page is up and running!');
    });
    
    app.get('/user', function (req, res) {
      var check = Users[req.query.name];
      var info = (check) ? check : 'User does not exist...';
      res.send(info);
    });

    app.get('/queryallcars', function (req, res) {
     
     
var Fabric_Client = require('fabric-client');
var path = require('path');
var util = require('util');
var os = require('os');

//
var fabric_client = new Fabric_Client();

// setup the fabric network
var channel = fabric_client.newChannel('mychannel');
var peer = fabric_client.newPeer('grpc://localhost:7051');
channel.addPeer(peer);

//
var member_user = null;
var store_path = path.join(__dirname, 'hfc-key-store');
console.log('Store path:'+store_path);
var tx_id = null;

// create the key value store as defined in the fabric-client/config/default.json 'key-value-store' setting
Fabric_Client.newDefaultKeyValueStore({ path: store_path
}).then((state_store) => {
	// assign the store to the fabric client
	fabric_client.setStateStore(state_store);
	var crypto_suite = Fabric_Client.newCryptoSuite();
	// use the same location for the state store (where the users' certificate are kept)
	// and the crypto store (where the users' keys are kept)
	var crypto_store = Fabric_Client.newCryptoKeyStore({path: store_path});
	crypto_suite.setCryptoKeyStore(crypto_store);
	fabric_client.setCryptoSuite(crypto_suite);

	// get the enrolled user from persistence, this user will sign all requests
	return fabric_client.getUserContext('user1', true);
}).then((user_from_store) => {
	if (user_from_store && user_from_store.isEnrolled()) {
		console.log('Successfully loaded user1 from persistence');
		member_user = user_from_store;
	} else {
		throw new Error('Failed to get user1.... run registerUser.js');
	}

	// queryCar chaincode function - requires 1 argument, ex: args: ['CAR4'],
	// queryAllCars chaincode function - requires no arguments , ex: args: [''],
	const request = {
		//targets : --- letting this default to the peers assigned to the channel
		chaincodeId: 'drug',
		fcn: 'queryAllDrugs',
		args: ['']
	};

	//const request = {
		//targets : --- letting this default to the peers assigned to the channel
	//	chaincodeId: 'drug',
		//fcn: 'queryDrug',
		//args: ['DRUG1']
	 // };

	// send the query proposal to the peer
	return channel.queryByChaincode(request);
}).then((query_responses) => {
	console.log("Query has completed, checking results");
	// query_responses could have more than one  results if there multiple peers were used as targets
	if (query_responses && query_responses.length == 1) {
		if (query_responses[0] instanceof Error) {
			console.error("error from query = ", query_responses[0]);
		} else {
			console.log("Response is ", query_responses[0].toString());
			res.send(query_responses[0].toString());
		}
	} else {		
		console.log("No payloads were returned from query");
	}
}).catch((err) => {
	console.error('Failed to query successfully :: ' + err);
});


    });
    
    // what port to run server on
    app.listen(3001, function () {
      console.log('server started on port 3001');
    });