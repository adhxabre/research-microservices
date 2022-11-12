const express = require('express')

const router = express.Router()

// Controller
const { getTodos, getTodo, addTodo, updateTodo, deleteTodo } = require('../controllers/todo')

// Route
router.get('/todos', getTodos)
router.get('/todo/:id', getTodo)
router.post('/todo', addTodo)
router.patch('/todo/:id', updateTodo)
router.delete('/todo/:id', deleteTodo)

module.exports = router