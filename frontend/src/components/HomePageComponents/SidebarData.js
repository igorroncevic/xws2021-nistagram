import React from 'react';
import * as FaIcons from 'react-icons/fa';
import * as AiIcons from 'react-icons/ai';
import * as IoIcons from 'react-icons/io';

export const SidebarData = [
    {
        title: 'Home',
        path: '/posts',
        icon: <AiIcons.AiFillHome />,
        cName: 'nav-text',
    },
    {
        title: 'Messages',
        path: '/chats',
        icon: <FaIcons.FaEnvelopeOpenText />,
        cName: 'nav-text'
    },
    {
        title: 'New post',
        path: '/newpost',
        icon: <FaIcons.FaEnvelopeOpenText />,
        cName: 'nav-text'
    },
    {
        title: 'Saved',
        path: '/saved',
        icon: <IoIcons.IoIosSave />,
        cName: 'nav-text'
    },
    {
        title: 'Profile',
        path: '/profile',
        icon: <IoIcons.IoMdPerson/>,
        cName: 'nav-text'
    },
    {
        title: 'Ovde dodavati sta sve treba vezano za usera kad se uloguje',
        //path: '/profile',
        //icon: <IoIcons.IoMdPeople />,
        //cName: 'nav-text'
    }


];