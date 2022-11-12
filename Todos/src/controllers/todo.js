let todos = [
    {
        id: 1,
        title: "Cuci tangan",
        isDone: true
    },
    {
        id: 2,
        title: "Jaga jarak",
        isDone: false
    },
    {
        id: 3,
        title: "Gunakan maskaer",
        idDone: true
    }
]

exports.getTodos = async (req, res) => {
    try {
        res.send({
            status: 'success',
            data: {
                todos
            }
        })
    } catch (error) {
        console.log(error)
        res.send({
            status: 'failed',
            message: 'Server Error'
        })
    }
}

exports.getTodo = async (req, res) => {
    try {
        const id = req.params.id
        const index = id - 1    
        res.send({
            status: 'success',
            data: {
                todo: todos[index]
            }
        })
    } catch (error) {
        console.log(error)
        res.send({
            status: 'failed',
            message: 'Server Error'
        })
    }
}

exports.addTodo = async (req, res) => {
    try {
        todos = [...todos, req.body];
        res.send({
            status: 'success',
            data: {
                todos
            }
        })
    } catch (error) {
        console.log(error)
        res.send({
            status: 'failed',
            message: 'Server Error'
        })
    }
}

exports.updateTodo = async (req, res) => {
    try {
        const { id } = req.params;
        todos[id - 1] = { ...todos[id - 1], ...req.body };
        
        res.send({
            status: 'success',
            data: {
                todo: todos[id - 1]
            }
        })
    } catch (error) {
        console.log(error)
        res.send({
            status: 'failed',
            message: 'Server Error'
        })
    }
}

exports.deleteTodo = async (req, res) => {
    try {
        const { id } = req.params;
        todos = todos.filter((todo) => todo.id != id);

        res.send({
            status: 'success',
            data: {
                todos
            }
        })
    } catch (error) {
        console.log(error)
        res.send({
            status: 'failed',
            message: 'Server Error'
        })
    }
}