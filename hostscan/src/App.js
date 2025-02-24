import React, { useState, useEffect,useContext } from 'react';
import { BrowserRouter as Router, Routes, Route,Navigate,useNavigate } from 'react-router-dom';
import Firstpage from './Pages/Firstpage';
import './App.css';
import LoginSignin from './Pages/LoginSignin';
import { DataContext, Passingdatas } from './Components/DataContext';
import Dashboard from './Pages/Dashboard';
import Scan from './Pages/Scan';
import Target from './Pages/Target';
import Risks from './Pages/Risks';
import Reports from './Pages/Reports';
import Integerations from './Pages/Integerations';
import Settings from './Pages/Settings';
import Support from './Pages/Support';
import APIDocs from './Pages/APIDocs';
import HelpCenter from './Pages/HelpCenter';
import Upgradeplan from './Pages/Upgradeplan';
import axios from "axios";

const getCookie = (name) => {
  const value = `; ${document.cookie}`;
  const parts = value.split(`; ${name}=`);
  if (parts.length === 2) return parts.pop().split(";").shift();
  return null;
};

function App() {
  var { emailId,setEmailId,userId, setUserId } = useContext(DataContext);  

  const [isAuthenticated, setIsAuthenticated] = useState(!!localStorage.getItem("email"));
  // const history = useNavigate();

  useEffect(() => {
    let emailFromCookie = getCookie("email");
    const jwtFromCookie = getCookie("jwt");

    emailFromCookie = decodeURIComponent(emailFromCookie).replace(/^j:/, '');
    emailFromCookie = JSON.parse(emailFromCookie)

    // Clear all cookies after storing email and setting the header
    const clearAllCookies = () => {
      document.cookie.split(';').forEach(function (c) {
        document.cookie = c.trim().split('=')[0] + '=; expires=Thu, 01 Jan 1970 00:00:00 GMT; path=/';
      });
    };
    
    if (emailFromCookie) {
      localStorage.setItem("email", emailFromCookie.email);
      localStorage.setItem("user_id", emailFromCookie.uid);
    }

    if (jwtFromCookie) {
      axios.defaults.headers.common["Authorization"] = jwtFromCookie;
      setIsAuthenticated(true);
    }else if (!(axios.defaults.headers.common["Authorization"])){
      console.log('****noo Authorization***');
      clearAllCookies();
      // history.push('/');  // Assuming '/' is the route you want to "reload"
      
    }

      // console.log('Authorization==> ',axios.defaults.headers.common["Authorization"]);
     

      
      
    
  }, []);



  // useEffect(() => {
  //   const email = localStorage.getItem("email");
  //   const user_id = localStorage.getItem("user_id");
  
  //   if (email && user_id) {
  //     console.log("inside useEffect - setting state:", email, user_id);
  //     setEmailId(email);
  //     setUserId(user_id);
  //   }
  // }, []); // Runs once after initial render

  useEffect(() => {
    const fetchData = async () => {
      const email = localStorage.getItem("email");
      const user_id = localStorage.getItem("user_id");
  
      if (email && user_id) {
        // console.log("inside useEffect - setting state:", email, user_id);
        setEmailId(email);
        setUserId(user_id);
      }
    };
  
    fetchData(); // ✅ Call the async function inside useEffect
  }, []);
  
  
  
    // // ✅ Log after state updates
    // useEffect(() => {
    //   console.log("Updated emailId:", emailId);
    //   // setEmailId(emailId);
    //   console.log("Updated userId:", userId);
    //   // setUserId(userId);
    // }, [emailId, userId]); // Runs when emailId or userId updates

  return (
    <div className="App">
      <Passingdatas>

        {!isAuthenticated ? (
          <Router>
            <Routes>
              <Route path='/' element={<Firstpage />} />
              <Route path='/login' element={<LoginSignin />} />
              <Route path='/sign' element={<LoginSignin />} />
              <Route path="*" element={<Navigate to="/" />} />
            </Routes>
          </Router>
        ) : (
          <Router>
            <Routes>
              <Route path='/' element={<Dashboard />} />
              <Route path='/dashboard' element={<Dashboard />} />
              <Route path='/Scannn' element={<Scan />} />
              <Route path='/Targettt' element={<Target />} />
              <Route path='/Risksss' element={<Risks />} />
              <Route path='/Reportsss' element={<Reports />} />
              <Route path='/Integerationsss' element={<Integerations />} />
              <Route path='/Settingsss' element={<Settings />} />
              <Route path='/Supporttt' element={<Support />} />
              <Route path='/APIDocsss' element={<APIDocs />} />
              <Route path='/HelpCenterrr' element={<HelpCenter />} />
              <Route path='/Upgradeplannn' element={<Upgradeplan />} />
              <Route path="*" element={<Navigate to="/" />} />
            </Routes>
          </Router>)}
      </Passingdatas>
    </div>
  );
}

export default App;
