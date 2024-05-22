import React from 'react'
import Loading from './Loading'
import Error from './Error'

export default function LoadingTodo(taskCache, tasksResponse) {
 if (taskCache === null && tasksResponse.isError) {
  return (<main>
   <Error />
  </main>)
 }
 if (taskCache === null && tasksResponse.isLoading) {
  return (<main>
   <Loading />
  </main>)
 }
 if (taskCache === null && tasksResponse.isSuccess) {
  taskCache = tasksResponse.data
 }
 if (taskCache === null) {
  return (<main>
   <Error />
  </main>)
 }
}

