import { useAuth } from './AuthProvider'
import { createContext, useContext, useEffect, useRef } from 'react'

const WebSocketContext = createContext()

const WebSocketProvider = ({children}) => {
    const socket = useRef(null)
    const token = useAuth().token
    const callbacks = useRef({})
    const userContext = useAuth()

    useEffect(() => {
        const ws = new WebSocket(`http://localhost:8080/api/v1/ws/chat?token=${token}`)
        socket.current = ws
    
        ws.onopen = () => {
            console.log('websocket connected', Date.now())
        }

        ws.onmessage = (event) => {
            const reqMessage = JSON.parse(event.data)
            console.log(reqMessage)
            const fromUser = reqMessage.from.user
            const message = reqMessage.message
            if(reqMessage.type === 2) {
                callbacks.current["explore-chat"](fromUser.username, false)
            }
            if(reqMessage.type === 1) {
                callbacks.current["receive-message"](fromUser.username, {
                    value: message,
                    send: 2
                })
            }
            if(reqMessage.type === 4) {
                callbacks.current["receive-friend-request"]()
            }
            if(reqMessage.type === 5) {
                callbacks.current["explore-chat"](fromUser.username, true)
                callbacks.current["add-friend"](fromUser)
            }
        }

        ws.onclose = (event) => {
            console.log(event)    
            // if(event.code===HttpStatusCode.Unauthorized) {
            //     userContext.LogoutAction()
            // } else {
            //     navigate("/")
            // }
            userContext.LogoutAction()
        }

        return () => {
            console.log("disconnect", Date.now())
            ws.close()
        }
    })

    const SendRequest = (req) => {
        if(socket.current && socket.current.readyState === WebSocket.OPEN) {
            socket.current.send(req)
        } else {
            console.log('websocket is not open')
        }
    }

    const RegisterCallback = (name, cb) => {
        if(!callbacks.current[name]) {
            callbacks.current[name] = cb
            console.log("callback registered")
        }
    }

    return (
        <WebSocketContext.Provider value={{SendRequest, RegisterCallback}}>
            {children}
        </WebSocketContext.Provider>
    )
}

export default WebSocketProvider

export const useWebSocket = () => useContext(WebSocketContext)