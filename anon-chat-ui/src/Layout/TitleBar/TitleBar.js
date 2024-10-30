import Title from './Title/Title';
import './TitleBar.css'
import NavBar from './NavBar/NavBar';

const TitleBar = ({onSwitch}) => {
    return(
        <div className="title-bar">
            <div className='title-container'>
                <Title/>
            </div>
            <NavBar onSwitch={onSwitch}/>
        </div>
    );
}

export default TitleBar;