'use strict';

const request = require('request');

function AuthService(url){
  this.url = url;
}

AuthService.prototype.login = function (email, password, cb) {
  request({
    url: this.url+'/auth?grant_type=password',
    method: 'POST',
    headers: {
      'Accept': 'application/json',
      'Content-Type': 'application/json'
    },
    body: {email: email, password: password},
    json: true
  }, (err, response, body) => {
    if(err || response.statusCode != '200'){
      cb("Invalid Username or Password", null);
    }else{
      cb(null, body)
    }
  });
};

module.exports = AuthService;
