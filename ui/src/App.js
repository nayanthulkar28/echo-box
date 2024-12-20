import {createBrowserRouter, RouterProvider} from 'react-router-dom'
import './App.css';
import Home from './Pages/Home';
import Login from './Pages/Login';
import ProtectedRoutes from './Components/ProtectedRoutes/ProtectedRoutes';
import NotFound from './Components/NotFound/NotFound';
import AuthProvider from './hooks/AuthProvider';
import SystemCompatible from './Components/SystemCompatible/SystemCompatible';

function App() {
  const router = createBrowserRouter([
    {
      path: "/login",
      element: <Login/>
    },
    {
      element: <ProtectedRoutes/>,
      children: [
        {
          path: "/",
          element: <Home/>
        }
      ]
    },
    {
      path: "/*",
      element: <NotFound/>
    }
  ])
  return (
    <div className="app">
      <AuthProvider>
        <SystemCompatible>
          <RouterProvider router={router}/>
        </SystemCompatible>
      </AuthProvider>
    </div>
  );
}

export default App;
