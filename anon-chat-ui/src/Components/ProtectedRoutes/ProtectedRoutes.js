import { Navigate, Outlet } from "react-router-dom"
import { useAuth } from "../../hooks/AuthProvider"
import WebSocketProvider from "../../hooks/Websocket"

const ProtectedRoutes = () => {
    const userContext = useAuth()
    if(userContext.user == null) return <Navigate to="/login" replace/>
    return (
        <WebSocketProvider> 
            <Outlet/>
        </WebSocketProvider>
    )
}

export default ProtectedRoutes