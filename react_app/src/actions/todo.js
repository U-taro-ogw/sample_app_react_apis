import * as todoApi from '../api/todo';

function fetchTasksSucceeded(tasks) {
  return {
    type: 'FETCH_TASKS_SUCCEEDED',
    payload: {
      tasks,
    },
  };
}

function fetchTasksFailed(error) {
  return {
    type: 'FETCH_TASKS_FAILED',
    payload: {
      error,
    },
  };
}

function fetchTasksStarted() {
  return {
    type: 'FETCH_TASKS_STARTED',
  };
}

export function fetchTasks() {
  return dispatch => {
    dispatch(fetchTasksStarted());

    todoApi
      .fetchTasks()
      .then(resp => {
        dispatch(fetchTasksSucceeded(resp.data.todos));
      })
      .catch(err => {
        dispatch(fetchTasksFailed(err.message));
      });
  };
}

function createTaskSucceeded(task) {
  return {
    type: 'CREATE_TASK_SUCCEEDED',
    payload: {
      task,
    },
  };
}

export function createTask({ title, description, status = 'Unstarted' }) {
  return dispatch => {
    todoApi.createTask({ title, description, status }).then(resp => {
      dispatch(createTaskSucceeded(resp.data));
    });
  };
}

function editTaskSucceeded(task) {
  return {
    type: 'EDIT_TASK_SUCCEEDED',
    payload: {
      task,
    },
  };
}

export function editTask(id, params = {}) {
  return (dispatch, getState) => {
    const task = getTaskById(getState().tasks.tasks, id);
    const updatedTask = Object.assign({}, task, params);
    todoApi.editTask(id, updatedTask).then(resp => {
      dispatch(editTaskSucceeded(resp.data.todo));
    });
  };
}


function getTaskById(tasks, id) {
  return tasks.find(task => task.id === id);
}


function deleteTaskSucceeded(id) {
  return {
    type: 'DELETE_TASK_SUCCEEDED',
    payload: {
      id,
    },
  };
}

export function deleteTask(id) {
  return (dispatch, getState) => {
    todoApi.deleteTask(id).then(resp => {
      console.log(resp)
      dispatch(deleteTaskSucceeded(id));
    });
  };
}
