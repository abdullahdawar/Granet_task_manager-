import type { Task, TaskInput } from '../types/task'

const API_URL = 'http://localhost:8080/api/tasks'

export async function getTasks(): Promise<Task[]> {
  const response = await fetch(API_URL)

  if (!response.ok) {
    throw new Error('Failed to load tasks')
  }

  return response.json()
}

export async function createTask(task: TaskInput): Promise<Task> {
  const response = await fetch(API_URL, {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    body: JSON.stringify(task),
  })

  if (!response.ok) {
    throw new Error('Failed to create task')
  }

  return response.json()
}

export async function updateTask(
  id: number,
  task: TaskInput,
): Promise<Task> {
  const response = await fetch(`${API_URL}/${id}`, {
    method: 'PUT',
    headers: {
      'Content-Type': 'application/json',
    },
    body: JSON.stringify(task),
  })

  if (!response.ok) {
    throw new Error('Failed to update task')
  }

  return response.json()
}

export async function deleteTask(id: number): Promise<void> {
  const response = await fetch(`${API_URL}/${id}`, {
    method: 'DELETE',
  })

  if (!response.ok) {
    throw new Error('Failed to delete task')
  }
}
