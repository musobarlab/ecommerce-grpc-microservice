'use strict';

const express = require('express');
const router = express.Router();

const AuthService = require('../service/auth');
const MeService = require('../service/me');

const TokenStorage = require('../utils/token_storage');


/* GET home page. */
router.get('/', (req, res, next) => {
  const BASE_URL = process.env.ORDER_BASE_URL;

  const tokenStorage = new TokenStorage();
  let token = tokenStorage.get();

  const meService = new MeService(BASE_URL);
  meService.me(token, (err, body) => {
    if(err){
      return res.render('pages/index', { title: 'Go Ecommerce', subtitle: 'Welcome', me: null });
    }

    let profile = JSON.parse(body);
    return res.render('pages/index', { title: 'Go Ecommerce', subtitle: 'Welcome', me: profile });
  });
});

router.get('/login', (req, res, next) => {
  const BASE_URL = process.env.ORDER_BASE_URL;

  const tokenStorage = new TokenStorage();
  let token = tokenStorage.get();

  const meService = new MeService(BASE_URL);
  meService.me(token, (err, body) => {
    if(err){
      return res.render('pages/login', { title: 'Go Ecommerce', subtitle: 'Login', me: null });
    }

    let profile = JSON.parse(body);
    return res.render('pages/login', { title: 'Go Ecommerce', subtitle: 'Login', me: profile });
  });
});

router.post('/auth', (req, res, next) => {
  const BASE_URL = process.env.AUTH_BASE_URL;
  let email = req.body.email;
  let password = req.body.password;

  if(email == '' || password == '') {
    return res.redirect('/');
  }

  const authService = new AuthService(BASE_URL);
  authService.login(email, password, (err, body) => {
    if(err){
      return res.redirect('/');
    }

    const tokenStorage = new TokenStorage();
    tokenStorage.store(body.accessToken);

    return res.redirect('/me');
  });
});

router.get('/me', (req, res, next) => {
  const BASE_URL = process.env.ORDER_BASE_URL;

  const tokenStorage = new TokenStorage();
  let token = tokenStorage.get();

  const meService = new MeService(BASE_URL);
  meService.me(token, (err, body) => {
    if(err){
      return res.redirect('/');
    }

    let profile = JSON.parse(body);
    return res.render('pages/me', { title: 'Go Ecommerce', subtitle: 'Customer Profile', me: profile });
  });
});

module.exports = router;
