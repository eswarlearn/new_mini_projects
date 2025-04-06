import React, { useContext,useState } from 'react'
import { RxDashboard } from "react-icons/rx";
import { LuServer } from "react-icons/lu";
import { TbActivityHeartbeat } from "react-icons/tb";
import { SlFlag } from "react-icons/sl";
import { VscBook } from "react-icons/vsc";
import { LuTrendingUpDown } from "react-icons/lu";
import { IoSettingsOutline } from "react-icons/io5";
import { CiMail } from "react-icons/ci";
import { LuCodeXml } from "react-icons/lu";
import { PiBookOpenThin } from "react-icons/pi";
import { MdOutlineRocket } from "react-icons/md";
import { IoIosArrowBack } from "react-icons/io";
import { IoIosArrowForward } from "react-icons/io";
import { DataContext } from './DataContext';
import { FaPlus } from "react-icons/fa6";
import { useNavigate } from 'react-router-dom';

const DashNav = () => {
  var nextpage=useNavigate()
  var {hiddencontent,setHiddenContent}=useContext(DataContext)
  var {bgnum,setbgnum}=useContext(DataContext)

  var clickscan=()=>{
    setbgnum(2)
    setHiddenContent(false);
    nextpage('/Scannn')
  }
  var clicktarget=()=>{
    setbgnum(1)
    setHiddenContent(false);
    nextpage('/Targettt')
  }
  var clickdash=()=>{
    setbgnum(0)
    setHiddenContent(false);
    nextpage('/')
  }
  var clickrisks=()=>{
    setbgnum(3)
    setHiddenContent(false);
    nextpage('/Risksss')
  }
  var clickreports=()=>{
    setbgnum(4)
    setHiddenContent(false);
    nextpage('/Reportsss')
  }
  var clickintegerations=()=>{
    setbgnum(5)
    setHiddenContent(false);
    nextpage('/Integerationsss')
  }
  var clicksettings=()=>{
    setbgnum(6)
    setHiddenContent(false);
    nextpage('/Settingsss')
  }
  var clicksupport=()=>{
    setbgnum(7)
    setHiddenContent(false);
    nextpage('/Supporttt')
  }
  var clickapidocs=()=>{
    setbgnum(8)
    setHiddenContent(false);
    nextpage('/APIDocsss')
  }
  var clickhelpcenter=()=>{
    setbgnum(9)
    setHiddenContent(false);
    nextpage('/HelpCenterrr')
  }
  var clickupgradeplan=()=>{
    setbgnum(10)
    setHiddenContent(false);
    nextpage('/Upgradeplannn')
  }
  
  return (
    <div className={`w-full  h-full ${hiddencontent ? 'md:max-w-[230px]' : 'md:max-w-[90px]'} mt-10 md:mt-[59px]`}>
    <div className={`flex w-full ${hiddencontent ? "translate-x-0 max-md:fixed" : "-translate-x-60"} transition-transform duration:200 md:duration-100 ease-in-out md:translate-x-0`} 
      style={{backgroundColor:'whitesmoke'}}>
      <div className={`relative z-[9999] flex flex-col max-md:w-full md:w-[200px] text-left p-4 gap-1 sm:mt-0`}>
        <button className='bg-black text-white px-5 py-4 items-center justify-center rounded-[4px] md:hidden block flex gap-2 w-auto md:w-max'>
          <FaPlus className='mt-1'/>
          <span className='text-[15px]'>New Scan</span>
        </button>
        {[
          { icon: <RxDashboard />, text: "Dashboard" ,onClick:clickdash },
          { icon: <LuServer />, text: "Targets" ,onClick:clicktarget },
          { icon: <TbActivityHeartbeat />, text: "Scans" ,onClick:clickscan },
          { icon: <SlFlag />, text: "Risks" ,onClick:clickrisks },
          { icon: <VscBook />, text: "Reports" ,onClick:clickreports },
          { icon: <LuTrendingUpDown />, text: "Integrations" ,onClick:clickintegerations },
          { icon: <IoSettingsOutline />, text: "Settings" ,onClick:clicksettings },
          { icon: <CiMail />, text: "Support" ,onClick:clicksupport },
          { icon: <LuCodeXml />, text: "API Docs" ,onClick:clickapidocs },
          { icon: <PiBookOpenThin />, text: "Help Center" ,onClick:clickhelpcenter },
          { icon: <MdOutlineRocket />, text: "Upgrade Plan", onClick:clickupgradeplan, special:true },
        ].map((item, index) => (
          <button
            key={index}
            className={`relative flex gap-3 rounded-[4px] ${hiddencontent ? 'max-w-full' : 'max-w-[40px]'} hover:max-w-full border border-transparent
              ${index==bgnum?'bg-white':index==10?'bg-[springgreen]':'bg-transparent'} ${index==7?'mt-20 md:mt-32':''} ${index!=10?'hover:border-black':'hover:border-green-500'} 
              px-2 py-[8px] group ${item.special ? `py-3 px-2 items-center justify-center border-2 border-green-600 text-black` : hiddencontent ? '':'hover:bg-white'} 
              transition-all duration-100 flex-nowrap z-[9999] `} onClick={item.onClick}
          >
            <span className={`mt-1 text-lg `}>{item.icon}</span>
            <span className={`left-10 z-[9999] ${hiddencontent ? 
              "translate-x-0" : "-translate-x-60"} overflow-hidden transition-transform duration-100 whitespace-nowrap ease-in-out 
              group-hover:translate-x-0 group-hover:block`}>
              {item.text}
            </span>
          </button>
        ))}
      </div>   
      <div className={`absolute ${hiddencontent ? 'ml-48' : 'ml-[70px]'} mt-7 whitespace-normal border-gray-300 border bg-white
       h-10 p-2 w-10 rounded-3xl hidden md:block cursor-pointer `} onClick={()=>setHiddenContent(!hiddencontent)}>
        <IoIosArrowBack className={`text-xl transition-transform duration-100 ${hiddencontent ? "block" : "hidden"}`}/>
        <IoIosArrowForward className={`text-xl ${hiddencontent ? "hidden" : "block"}`}/>
      </div>   
    </div>
   
    </div>
  )
}

export default DashNav
