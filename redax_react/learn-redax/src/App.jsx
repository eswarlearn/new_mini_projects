import { useState } from 'react'
import reactLogo from './assets/react.svg'
import viteLogo from '/vite.svg'
import './App.css'
import ComponentAdd from './ComponentAdd.jsx'
import CustomerView from './CustomerView.jsx'
import { Provider } from 'react-redux'
import { store } from './Store.jsx'

function App() {
 

  return (
    <Provider store={store}>
      <h3>Redux Customer Add Component</h3>
     <ComponentAdd/>
     <CustomerView/>
    </Provider>
  )
}

export default App
