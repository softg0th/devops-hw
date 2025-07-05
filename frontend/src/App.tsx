import React from 'react';
import './App.css';
import "./pages/AnimalsPageController";
import {AnimalsPageContainer} from "./pages/AnimalsPageController";
import {Provider} from "react-redux";
import {store} from "./api/store";

function App() {
    return (
        <div className="App">
            <div className="content">
                <Provider store={ store }>
                    <AnimalsPageContainer/>
                </Provider>
            </div>
        </div>
    );
}

export default App;
