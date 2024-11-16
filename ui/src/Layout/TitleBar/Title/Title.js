import './Title.css'

const Title = ({class_name}) => {
    return(
        <div className={"title " + class_name}>
            Echo Box
        </div>
    );
}

export default Title;