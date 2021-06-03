import React, {useEffect, useState} from 'react';
import * as FaIcons from 'react-icons/fa';
import * as AiIcons from 'react-icons/ai';
import { Link,useHistory } from 'react-router-dom';
import '../../style/HomePage.css';
import { IconContext } from 'react-icons';
import {SidebarData} from "./SidebarData";
import {Button} from "react-bootstrap";


function HomePage() {
    const [sidebar, setSidebar] = useState(false);

    const showSidebar = () => setSidebar(!sidebar);

    const history = useHistory()

    useEffect(() => {
        document.body.style.backgroundColor = "#ffeecc"
    });

    function  logOut(){
        console.log("BLA")
        history.push('/');
    }

    return (
        <>
            <IconContext.Provider value={{color: '#fff'}}>
                <div className='navbar'>
                    <Link to='#' className='menu-bars'>
                        <FaIcons.FaBars onClick={showSidebar}/>
                    </Link>
                    <p className='text'>Nistagram</p>
                    <Button style={{background: '#ffeecc', borderColor: '#ffeecc', color: '#ff4d4d'}} onClick={logOut} >Log out</Button>

                </div>
                <nav className={sidebar ? 'nav-menu active' : 'nav-menu'}>
                    <ul className='nav-menu-items' onClick={showSidebar}>
                        <li className='navbar-toggle'>
                            <Link to='#' className='menu-bars'>
                                <AiIcons.AiOutlineClose/>
                            </Link>
                        </li>
                        {SidebarData.map((item, index) => {
                            return (
                                <li key={index} className={item.cName}>
                                    <Link to={item.path}>
                                        {item.icon}
                                        <span>{item.title}</span>
                                    </Link>
                                </li>
                            );
                        })}
                    </ul>
                </nav>
            </IconContext.Provider>
        </>
    );
}

export default HomePage;