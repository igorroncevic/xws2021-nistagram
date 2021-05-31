import  React from "react";
import {LoginPage} from './pages/LoginPage.js'

function App () {
    return (
        <div className="App">
            Here is the Latest React version: <strong>{React.version}</strong>
        <LoginPage/>
        </div>
    );
}
export default App