const express = require('express')

const router = express.Router()

// Controller
// import controller here
const { addUser } = require('../controllers/user')

// Route
// add route here
router.post('/user', addUser)

module.exports = router