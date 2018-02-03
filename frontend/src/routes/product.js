'use strict';

const express = require('express');
const router = express.Router();

const ProductService = require('../service/product');
const CategoryService = require('../service/category');

const TokenStorage = require('../utils/token_storage');

router.get('/products', (req, res, next) => {
  const BASE_URL = process.env.ORDER_BASE_URL;

  const tokenStorage = new TokenStorage();
  let token = tokenStorage.get();

  const meService = new MeService(BASE_URL);
  const productService = new ProductService(BASE_URL);
  meService.me(token, (err, body) => {
    if(err){
      productService.getAll((err, products) => {
        if(err){
          console.log(err);
          return res.render('pages/index', { title: 'Go Ecommerce', subtitle: 'Welcome', products: null, me: null });
        }
        let productJson = JSON.parse(products);
        return res.render('pages/index', { title: 'Go Ecommerce', subtitle: 'Welcome', products: productJson, me: null });
      });
    }

    productService.getAll((err, products) => {
      if(err){
        return console.log(err);
      }

      let profile = JSON.parse(body);
      let productJson = JSON.parse(products);
      return res.render('pages/index', { title: 'Go Ecommerce', subtitle: 'Welcome', products: productJson, me: profile });
    });
  });
});
