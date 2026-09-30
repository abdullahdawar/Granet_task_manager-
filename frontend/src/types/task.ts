export interface Task {
  id: number
  title: string
  description: string
  due_date: string
  status: string
}

export interface TaskInput {
  title: string
  description: string
  due_date: string
  status: string
}
