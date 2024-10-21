// endpoint
const apiTodoListEndpoint = "http://localhost:8080/api/todos"
const apiDeleteEndpoint = "http://localhost:8080/api/delete"
const apiUpdateEndpoint = "http://localhost:8080/api/update"

// button text
const taskDoneButtonText = "未完了に戻す"
const taskNotDoneButtonText = "タスク完了"

fetch(apiTodoListEndpoint)
    .then(response => {
        if (!response.ok) {
            throw new Error("Network response was not ok " + response.statusText)
        }
        return response.json()
    })
    .then(data => {
        data.forEach(todo => {
            displayTodo(todo)
        });
    })
    .catch(error => {
        console.error("There was a problem with the fetch operation:", error)
    })

function displayTodo(todo) {
    const todoList = document.getElementById("todoList")

    // div要素を作成
    const todoItem = document.createElement("div")
    // classとしてtodoItemを追加
    todoItem.classList.add("todoItem")
    todoItem.setAttribute("data-id", todo.ID)

    // Contentの表示
    const contentElement = document.createElement("p")
    contentElement.textContent = todo.Content
    todoItem.appendChild(contentElement)

    // Doneボタンを表示
    const doneButton = createDoneButton(todo.ID, todo.Done)
    todoItem.appendChild(doneButton)

    // Deleteボタンの設置
    const deleteButton = createDeleteButton(todo.ID)
    todoItem.appendChild(deleteButton)

    // /Todoの要素をtodoListに追加
    todoList.appendChild(todoItem)
}


function createDeleteButton (id) {
    const deleteButton = document.createElement("button")
    deleteButton.textContent = "削除"
    deleteButton.onclick = function() {
        deleteTodoItem(id)
    }
    return deleteButton
}

function deleteTodoItem(id) {
    // リクエストする対象のIDをJSONに詰める
    const requestData = {
        ID: id
    }

    // /api/deleteに対して、リクエストを送る
    fetch(apiDeleteEndpoint, {
        method: "DELETE",
        headers: {
            "Content-Type": "application/json"
        },
        body: JSON.stringify(requestData)
    })
    .then((response) => {
        // responseがエラーの場合
        if (!response.ok) {
            throw new Error("Failed to delete Todo item")
        }
        return response.json()
    })
    .then(responseJson => {
        // responseで返ってきたJSONの内容により処理を分岐
        // 成功の場合
        if (responseJson.result === "SUCCESS") {
            removeTodoElement(id)
        // 失敗の場合
        } else {
            alert(`Error: ${responseJson.result}`)
        }
    })
    // フェッチそのもののエラーをキャッチする
    .catch(error => {
        console.error("There was a problem with the fetch operation:", error)
    })

}

function removeTodoElement(id) {
    const todoItem = document.querySelector(`[data-id='${id}']`)
    if (todoItem) {
        todoItem.remove()
    }
}

function createDoneButton (id, isDone) {
    const doneButton = document.createElement("button")
    doneButton.textContent = isDone ? taskDoneButtonText : taskNotDoneButtonText 
    doneButton.onclick = function() {
        updateDoneButton(id)
    }

    return doneButton
}

function updateDoneButton(id) {
    // リクエストする対象のIDをJSONに詰める
    const requestData = {
        ID: id
    }

    // /api/deleteに対して、リクエストを送る
    fetch(apiUpdateEndpoint, {
        method: "PUT",
        headers: {
            "Content-Type": "application/json"
        },
        body: JSON.stringify(requestData)
    })
    .then((response) => {
        // responseがエラーの場合
        if (!response.ok) {
            throw new Error("Failed to update Todo item")
        }
        return response.json()
    })
    .then(responseJson => {
        // responseで返ってきたJSONの内容により処理を分岐
        // 成功の場合
        if (responseJson.result === "SUCCESS") {
            toggleUpdateButton(id)
        // 失敗の場合
        } else {
            alert(`Error: ${responseJson.result}`)
        }
    })
    // フェッチそのもののエラーをキャッチする
    .catch(error => {
        console.error("There was a problem with the fetch operation:", error)
    })
}

function toggleUpdateButton(id) {
    const todoItem = document.querySelector(`[data-id='${id}']`)
    if (todoItem) {
        const doneButton = todoItem.querySelector("button")

        if (doneButton) {
            if(doneButton.textContent === taskDoneButtonText) {
                doneButton.textContent = taskNotDoneButtonText
            } else if (doneButton.textContent === taskNotDoneButtonText) {
                doneButton.textContent = taskDoneButtonText
            }
        }
    } 
}

