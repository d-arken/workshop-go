import logo from './logo.svg';
import './App.css';
import {KindeProvider} from "@kinde-oss/kinde-auth-react";
import {Home} from "./Home";

function App() {
  return (
      <KindeProvider domain={"https://darkmdev.kinde.com"} clientId={"d8db0051dc7647c49da254d35c736f37"}
                     logoutUri={window.location.origin}
                     redirectUri={window.location.origin}>
        <div className="App">

            <Home />
        </div>


      </KindeProvider>
  )
}

export default App;
