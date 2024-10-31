import Title from './Title/Title';
import './TitleBar.css'
import NavBar from './NavBar/NavBar';

const TitleBar = ({onSwitch}) => {
    return(
        <div className="title-bar">
            <div className='title-container'>
                <Title class_name="title-home"/>
            </div>
            <NavBar onSwitch={onSwitch}/>
        </div>
    );
}

export default TitleBar;