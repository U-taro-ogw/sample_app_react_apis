import axios from 'axios';

const API_BASE_URL = 'http://localhost:8083';

const client = axios.create({
  baseURL: API_BASE_URL,
  headers: {
    'Content-Type': 'application/json'
  },
});

export function fetchTasks() {
  return client.get('/todos');
}

export function createTask(params) {
  console.log(params)
  return client.post('/todos', params);
}

export function editTask(id, params) {
  console.log(params)
  return client.put(`/todos/${id}`, params);
}

export function deleteTask(id) {
  return client.delete(`/todos/${id}`);
}
