document.getElementById('add-button').addEventListener('click', addTodo);

function addTodo() {
    const input = document.getElementById('todo-input');
    const todo = input.value;
    if (todo) {
        const li = document.createElement('li');
        li.textContent = todo;
        li.addEventListener('click', toggleComplete);
        document.getElementById('todo-list').appendChild(li);
        input.value = '';
    }
}

function toggleComplete(event) {
    event.target.classList.toggle('completed');
}
