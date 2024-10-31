import './Title.css'

const Title = ({class_name}) => {
    return(
        <div className={"title " + class_name}>
            Anon Chat
        </div>
    );
}

export default Title;