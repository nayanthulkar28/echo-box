import {createBrowserRouter, RouterProvider} from 'react-router-dom'
import './App.css';
import Home from './Pages/Home';
import Login from './Pages/Login';
import ProtectedRoutes from './Components/ProtectedRoutes/ProtectedRoutes';
import NotFound from './Components/NotFound/NotFound';
import AuthProvider from './hooks/AuthProvider';

function App() {
  const router = createBrowserRouter([
    {
      path: "echo-box/login",
      element: <Login/>
    },
    {
      element: <ProtectedRoutes/>,
      children: [
        {
          path: "echo-box/",
          element: <Home/>
        }
      ]
    },
    {
      path: "echo-box/*",
      element: <NotFound/>
    }
  ])
  return (
    <div className="app">
      <AuthProvider>
        <RouterProvider router={router}/>
      </AuthProvider>
    </div>
  );
}

export default App;
