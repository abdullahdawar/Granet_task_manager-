import { useEffect, useState } from 'react'
import './App.css'
import {
  createTask,
  deleteTask,
  getTasks,
  updateTask,
} from './services/taskApi'
import type { Task, TaskInput } from './types/task'

const emptyForm: TaskInput = {
  title: '',
  description: '',
  due_date: '',
  status: 'pending',
}

function App() {
  const [tasks, setTasks] = useState<Task[]>([])
  const [form, setForm] = useState<TaskInput>(emptyForm)
  const [editingId, setEditingId] = useState<number | null>(null)
  const [loading, setLoading] = useState(true)
  const [saving, setSaving] = useState(false)
  const [error, setError] = useState('')

  async function loadTasks() {
    try {
      setError('')
      const data = await getTasks()
      setTasks(data)
    } catch {
      setError('Could not load tasks.')
    } finally {
      setLoading(false)
    }
  }

  useEffect(() => {
    loadTasks()
  }, [])

  function handleChange(
    event: React.ChangeEvent<HTMLInputElement | HTMLTextAreaElement | HTMLSelectElement>,
  ) {
    const { name, value } = event.target
    setForm((current) => ({ ...current, [name]: value }))
  }

  async function handleSubmit(event: React.FormEvent) {
    event.preventDefault()

    if (!form.title.trim() || !form.due_date) {
      setError('Title and due date are required.')
      return
    }

    try {
      setSaving(true)
      setError('')

      if (editingId !== null) {
        await updateTask(editingId, form)
      } else {
        await createTask(form)
      }

      setForm(emptyForm)
      setEditingId(null)
      await loadTasks()
    } catch {
      setError('Could not save task.')
    } finally {
      setSaving(false)
    }
  }

  function startEditing(task: Task) {
    setEditingId(task.id)
    setForm({
      title: task.title,
      description: task.description,
      due_date: task.due_date.slice(0, 10),
      status: task.status,
    })
    setError('')
  }

  function cancelEditing() {
    setEditingId(null)
    setForm(emptyForm)
    setError('')
  }

  async function handleDelete(id: number) {
    try {
      setError('')
      await deleteTask(id)

      if (editingId === id) {
        cancelEditing()
      }

      await loadTasks()
    } catch {
      setError('Could not delete task.')
    }
  }

  return (
    <main className="app">
      <header className="header">
        <div>
          <p className="eyebrow">TASK MANAGEMENT</p>
          <h1>Stay on top of your work.</h1>
          <p className="subtitle">
            Create, organize and complete your tasks in one place.
          </p>
        </div>
        <div className="task-count">
          <strong>{tasks.length}</strong>
          <span>{tasks.length === 1 ? 'Task' : 'Tasks'}</span>
        </div>
      </header>

      {error && <div className="error-message">{error}</div>}

      <section className="layout">
        <form className="task-form" onSubmit={handleSubmit}>
          <h2>{editingId !== null ? 'Edit Task' : 'New Task'}</h2>

          <label>
            Title
            <input
              name="title"
              value={form.title}
              onChange={handleChange}
              placeholder="What needs to be done?"
              maxLength={255}
              required
            />
          </label>

          <label>
            Description
            <textarea
              name="description"
              value={form.description}
              onChange={handleChange}
              placeholder="Add some details..."
              rows={4}
            />
          </label>

          <label>
            Due Date
            <input
              type="date"
              name="due_date"
              value={form.due_date}
              onChange={handleChange}
              required
            />
          </label>

          <label>
            Status
            <select name="status" value={form.status} onChange={handleChange}>
              <option value="pending">Pending</option>
              <option value="in_progress">In Progress</option>
              <option value="completed">Completed</option>
            </select>
          </label>

          <div className="form-actions">
            <button className="primary-button" type="submit" disabled={saving}>
              {saving
                ? 'Saving...'
                : editingId !== null
                  ? 'Update Task'
                  : 'Add Task'}
            </button>

            {editingId !== null && (
              <button
                className="secondary-button"
                type="button"
                onClick={cancelEditing}
              >
                Cancel
              </button>
            )}
          </div>
        </form>

        <section className="tasks-section">
          <div className="section-heading">
            <h2>Your Tasks</h2>
            <span>{tasks.length} total</span>
          </div>

          {loading ? (
            <p className="empty-state">Loading tasks...</p>
          ) : tasks.length === 0 ? (
            <div className="empty-state">
              <h3>No tasks yet</h3>
              <p>Create your first task using the form.</p>
            </div>
          ) : (
            <div className="task-list">
              {tasks.map((task) => (
                <article className="task-card" key={task.id}>
                  <div className="task-card-top">
                    <h3>{task.title}</h3>
                    <span className={`status ${task.status}`}>
                      {task.status.replace('_', ' ')}
                    </span>
                  </div>

                  <p>{task.description || 'No description provided.'}</p>

                  <div className="task-meta">
                    <span>Due {task.due_date.slice(0, 10)}</span>
                    <div>
                      <button
                        className="text-button"
                        onClick={() => startEditing(task)}
                      >
                        Edit
                      </button>
                      <button
                        className="text-button danger"
                        onClick={() => handleDelete(task.id)}
                      >
                        Delete
                      </button>
                    </div>
                  </div>
                </article>
              ))}
            </div>
          )}
        </section>
      </section>
    </main>
  )
}

export default App
