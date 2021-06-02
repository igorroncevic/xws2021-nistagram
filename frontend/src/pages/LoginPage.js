import React, {useState} from 'react';


export function LoginPage(){
    return (
            <div>
                <div className="search-input">
                    <input type="text" placeholder="Search"/>
                </div>
                <h1 className="h1">Search Results</h1>
                <div className="books">
            </div>
        </div>
    );
}