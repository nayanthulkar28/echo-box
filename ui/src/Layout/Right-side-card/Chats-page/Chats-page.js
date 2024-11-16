import Chat from '../../../Components/Chat/Chat';
import './Chats-page.css'

const ChatsPage = ({friend, isFriend, messages, appendMessage}) => {
    return(
        <div className="chats-page">
            <Chat 
                friend={friend}
                isFriend={isFriend}
                messages={messages}
                appendMessage={appendMessage}/>
        </div>
    );
}

export default ChatsPage;