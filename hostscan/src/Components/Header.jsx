import React, { useState ,useEffect,useContext} from 'react'
import Iconimg from '../Images/download.png'
import { RxHamburgerMenu } from "react-icons/rx";
import Navcom from './Navcom';
import { IoIosAdd } from 'react-icons/io';
import { useNavigate } from 'react-router-dom';
import { DataContext } from './DataContext';

const Header = () => {
    var navigate=useNavigate()
    var [open,isOpen]=useState(false)
    var {inn,setInn}=useContext(DataContext)
    var Login=()=>{
        // isOpen(!open)
        setInn('in')
        navigate('/login')
    }   
    var Signin=()=>{
        // isOpen(!open)   
        setInn('up')     
        navigate('/sign')
    }
    useEffect(() => {
        const handleResize = () => {
          if (window.innerWidth >= 1024) {
            isOpen(false); 
          } 
        };
        window.addEventListener("resize", handleResize);
        handleResize();
        return () => window.removeEventListener("resize", handleResize);
    }, []);

  return (
    <div className='container' style={{backgroundColor:'whitesmoke'}}>
        <header className='fixed top-0 z-20 w-full flex justify-center border-b-[1.5px] border-gray-300 px-8 py-3 bg-white/150 backdrop-blur-md'>
            <nav className='flex w-full items-center justify-between max-w-[1400px]'>
                <div className='flex gap-3'>
                    <img src={Iconimg} className='h-6 w-6'/>
                    <h2 className='font-bold' style={{fontFamily:'inherit'}}>HostedScan</h2>
                </div>
                <div className={`w-full items-center justify-center lg:flex gap-1 hidden lg:block`}>
                    <Navcom clickeve={()=>Login()} signinn={()=>Signin()} />
                </div>
                <div className='gap-6 flex flex-row'>
                    <button onClick={()=>Login()} className='hidden sm:block border-white/50 bg-gray-200 text-black font-bold flex-none px-5 py-2 rounded-md'>Log in</button>
                    <button onClick={()=>Signin()} className='hidden sm:block whitespace-nowrap border bg-black text-[#f8fafc] font-bold flex-none px-5 py-2 rounded-md'>Sign up</button>
                    <RxHamburgerMenu className='font-bold text-3xl sm:mt-2 lg:hidden' onClick={()=>isOpen(!open)}/>                    
                </div>
            </nav>
        </header>
        <div className={`lg:hidden sm:py-20 px-5 py-16 w-full fixed`}>
            <div className={`w-full ${open ? 'block' : 'hidden'} text-left bg-slate-800 text-white 
                rounded-xl`}>
                <Navcom clickeve={()=>Login()} signinn={()=>Signin()}/>            
            </div>
        </div>
    </div>
  )
}

export default Header
