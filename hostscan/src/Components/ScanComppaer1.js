import React, { useContext, useState,useRef,useEffect } from 'react'
import { FaPlus } from "react-icons/fa6";
import { DataContext } from './DataContext';
import { IoFilterSharp } from "react-icons/io5";
import { SlArrowRight } from "react-icons/sl";
import { TfiReload } from "react-icons/tfi";
import { PiNotePencilLight } from "react-icons/pi";
import { RiDeleteBinLine } from "react-icons/ri";
import { SlArrowDown } from "react-icons/sl";
import { BsExclamationCircle } from "react-icons/bs";
import { TbReport } from "react-icons/tb";
import FilterTarget from './FilterTarget';


const ScanComppaer1 = () => {
    var modalRef=useRef(null)
    var {hiddencontent,setHiddenContent}=useContext(DataContext)
    var [scanDetails, setScanDetails]=useState(true)
    var {filteropen, setFilterOpen}=useContext(DataContext)
    var [selectedRow,setSelectedRow]=useState(null)
    var [selectedScheduleRow,setSelectedScheduleRow]=useState(null)
    var scanresult=[
        { id: 1, name: "OpenVAS", target:"www.example.com", frequency: "Overlimit", days: "3 days ago" },
        { id: 2, name: "OpenVAS", target:"https://www.codering.com/", frequency: "Overlimit", days: "8 days ago" },
        { id: 3, name: "OWASP ZAP", target:"https://www.codering.com/", frequency: "Overlimit", days: "8 days ago" },
        { id: 4, name: "Nmap", target:"https://www.codering.com/", frequency: "Overlimit", days: "8 days ago" },
        { id: 5, name: "OpenVAS", target:"https://www.codering.com/", frequency: "Overlimit", days: "8 days ago" },
        { id: 6, name: "OWASP ZAP", target:"https://www.codering.com/", frequency: "Overlimit", days: "8 days ago" },
        { id: 7, name: "Nmap", target:"https://www.codering.com/", frequency: "Overlimit", days: "8 days ago" },
        { id: 8, name: "OpenVAS", target:"https://www.codering.com/", frequency: "Report", days: "8 days ago" },
        { id: 9, name: "OWASP ZAP", target:"https://www.codering.com/", frequency: "Report", days: "8 days ago" },
        { id: 10, name: "Nmap", target:"https://www.codering.com/", frequency: "Report", days: "8 days ago" },
    ];
    var scanData = [
        { id: 1, name: "OWASP ZAP", frequency: "Monthly", days: "in 23 days" },
        { id: 2, name: "Nmap", frequency: "Weekly", days: "in 10 days" },
        { id: 3, name: "Burp Suite", frequency: "weekly", days: "in 15 days" },
        { id: 4, name: "OWASP ZAP", frequency: "Monthly", days: "in 23 days" },
        { id: 5, name: "Nmap", frequency: "Weekly", days: "in 10 days" },
        { id: 6, name: "Burp Suite", frequency: "weekly", days: "in 15 days" },
    ];
    var toggleAction=(e)=>{
        setSelectedRow(e == selectedRow ? null : e)
        setScanDetails(!scanDetails)
    }
    var toggleschedulescan=(e)=>{
        setSelectedScheduleRow(e==selectedScheduleRow ? null : e)
        setScanDetails(!scanDetails)
    }
    var scantext=[
        {id:1, name:'schedule', value:'MONTHLY'},
        {id:2, name:'id', value:'6790bb54ecc390814192d35c'},
        {id:3, name:'scan_type', value:'OPENVAS'},
        {id:4, name:'email_notification', value:'ALWAYS'},
        {id:5, name:'next_run', value:'2025-02-22T09:33:08.572Z'},
        {id:6, name:'created_at', value:'2025-01-22T09:33:08.573Z'},
    ]
    var scheduletext=[
        {id:1, name:'schedule', value:'MONTHLY'},
        {id:2, name:'id', value:'6790bb54ecc390814192d35c'},
        {id:3, name:'scan_type', value:'OPENVAS'},
        {id:4, name:'email_notification', value:'ALWAYS'},
        {id:5, name:'next_run', value:'2025-02-22T09:33:08.572Z'},
        {id:6, name:'created_at', value:'2025-01-22T09:33:08.573Z'},
    ]
    useEffect(() => {
        const handleClickOutside = (event) => {
          if (modalRef.current && !modalRef.current.contains(event.target)) {
            setFilterOpen(true)
          }
        };
        document.addEventListener("mousedown", handleClickOutside);
        return () => {
          document.removeEventListener("mousedown", handleClickOutside);
        };
      }, []);

  return (
    <div className={`mt-20 px-4 flex ${hiddencontent?'md:ml-[200px]':'md:ml-[90px]'} flex-col gap-8 md:mt-32 w-full`}>
        <div className='bg-orange-100 border-2 border-orange-300 text-sm text-orange-800 p-3 rounded-[7px] flex items-start px-4'>
            <span>Request exceeded the free scan credits limit. <span className='font-bold underline'>Upgrade your plan</span></span>
        </div>
        <div ref={modalRef} className={`fixed ${filteropen ? 'hidden' : 'block'} max-sm:w-full mt-[-25px] right-0 sm:right-4 border px-5 py-4 rounded-[7px] bg-white flex flex-col `}>
            <FilterTarget/>
        </div>
        <div className='gap-2 flex md:flex-row flex-col items-start justify-start md:items-center md:justify-between'>
            <div className=''>
                <h1 className='font-bold text-4xl'>Scans</h1>
            </div>
            <div className='flex gap-2'>
                <button className='bg-black sm:hidden block text-white px-5 py-3 items-center justify-center rounded-[4px] flex gap-2 w-auto md:w-max'>
                    <FaPlus className='mt-1'/>
                    <span className='text-[15px] font-bold'>New Scan</span>
                </button>
                <button onClick={()=>setFilterOpen(!filteropen)} className='border items-center px-4 py-3 justify-center rounded-[4px] flex gap-2 w-auto md:w-max'
                    style={{backgroundColor:'whitesmoke'}}>
                    <IoFilterSharp className='mt-1'/>
                    <span className='text-[15px]'>Filters</span>
                </button>
            </div>
        </div>
        <div className='border-2 hidden border-gray-300 rounded-lg p-8 border-dashed'>
            <span className='text-gray-500 text-sm'>No data to display</span>
        </div>
        <div className='mt-12 flex flex-col'>                
            <div className="space-y-0">
                {/* First Row */}
                <div className="flex w-full items-center text-xs font-medium justify-between">
                    <div className="w-[180px] flex-shrink-0">
                        <a className="block">Scan</a>
                    </div>
                    <div className="grid grid-cols-4 md:grid-cols-5 flex-1 text-center">
                        <a>Target(s)</a>
                        <a></a>
                        <a className='text-start'>Progress/Results</a>
                        <a className="hidden md:block">Created</a>
                        <a></a>
                    </div>
                </div>
                <div>
                    <hr className='w-full mt-4'/>
                </div>                   
                <div className="space-y-0">
                    {scanresult.map((scan, index) => (
                        <div className='flex flex-col gap-0'>
                            <div key={scan.id} onClick={() => toggleAction(scan.id)} className="flex w-full hover:bg-gray-100 items-center text-xs font-medium justify-between">
                                <div className="w-[180px] flex-shrink-0 flex items-center">
                                    <div className="hover:bg-gray-300 px-[18px] py-4 flex items-center justify-center">
                                        <button>
                                            {selectedRow == scan.id ? (
                                                <SlArrowDown/>
                                            ):(
                                                <SlArrowRight/>
                                            )}
                                        </button>
                                    </div>
                                    <span className="ml-2">{scan.name}</span>
                                </div>
                                <div className="grid grid-cols-4 md:grid-cols-5 flex-1 items-center text-center">
                                    <a className=''>{scan.target}</a>
                                    <a></a>
                                    <a className='flex gap-2'>{index < 7 ? (       
                                        <>                                
                                            <span className='text-black'>{scan.frequency}</span>
                                            <BsExclamationCircle className='text-red-500'/> 
                                        </>                                       
                                    ):(
                                        <div className='flex gap-3 items-center'>
                                            <div className='flex'>
                                                <TbReport className='text-blue-700 text-lg'/>
                                                <span className='text-blue-700'>{scan.frequency}</span>
                                            </div>
                                            <div className="h-4 border border-gray-300"></div>
                                            <div className='flex gap-1'>
                                                <SlArrowRight className='text-[8px] mt-1'/>
                                                <span>Output</span>
                                            </div>
                                        </div>
                                    )}</a>
                                    <a className="hidden md:block">{scan.days}</a>
                                    <a className="flex items-end gap-2 xl:ml-40">
                                        <div className="border-[1px] hover:text-blue-400 text-[16px] border-blue-300 p-2 rounded-md"><TfiReload /></div>
                                    </a>
                                </div>
                            </div>
                            <div>
                                <hr className='w-full'/>
                            </div>
                            {selectedRow == scan.id ? (
                                scantext.map((txt,index)=>(
                                    <div className='flex mt-2 gap-4'>
                                        <div className='flex flex-col items-end'>
                                            <hr className='w-[150px]'/>    
                                            <span className='md:text-xs text-[10px] text-gray-500 mr-3 mt-2'>{txt.name}</span>
                                        </div> 
                                        <div className='flex flex-col items-start'>   
                                            <span className={`md:text-xs text-[10px] font-bold mt-2 ${index === scheduletext.length - 1 ? 'mb-5' : ''}`}>
                                                {txt.value}
                                            </span>
                                        </div> 
                                    </div> 
                                ))
                            ):(
                                <span></span>
                            )}
                        </div>
                    ))}
                </div>
            </div>
        </div>
        <div className='gap-2 mt-8 w-full flex flex-col'>
            <div className='flex ml-4 items-start justify-start'>
                <h2 className='font-bold text-2xl'>Scheduled Scans</h2>
            </div>
            <div className='mt-12 flex flex-col'>                
                <div className="space-y-0">
                    {/* First Row */}
                    <div className="flex w-full items-center text-xs font-medium justify-between">
                        <div className="w-[180px] flex-shrink-0">
                            <a className="block">Target(s)</a>
                        </div>
                        <div className="grid grid-cols-4 md:grid-cols-5 flex-1 text-center">
                            <a>Scan</a>
                            <a>Schedule</a>
                            <a className="hidden md:block">Last ran</a>
                            <a className="hidden md:block">Next run</a>
                        </div>
                    </div>
                    <div>
                        <hr className='w-full mt-4'/>
                    </div>                   
                    <div className="space-y-0">
                        {scanData.map((scan, index) => (
                            <div className='flex flex-col gap-0'>
                                <div key={scan.id} onClick={() => toggleschedulescan(scan.id)} className="flex w-full items-center text-xs font-medium hover:bg-gray-100 justify-between">
                                    <div className="w-[180px] flex-shrink-0 flex items-center">
                                        <div className="hover:bg-gray-300 px-[18px] py-4 flex items-center justify-center">
                                            <button>
                                                {selectedScheduleRow == scan.id ? (
                                                    <SlArrowDown/>
                                                ):(
                                                    <div>
                                                        <SlArrowRight/>
                                                    </div>
                                                )}
                                            </button>
                                        </div>
                                        <span className="ml-2">(all targets)</span>
                                    </div>
                                    <div className="grid grid-cols-4 md:grid-cols-5 flex-1 items-center text-center">
                                        <a>{scan.name}</a>
                                        <a>{scan.frequency}</a>
                                        <a className="hidden md:block"></a>
                                        <a className="hidden md:block">{scan.days}</a>
                                        <a className="flex items-start gap-2 xl:ml-20">
                                            <div className="border-[1px] text-[16px] hover:text-blue-400 border-blue-300 p-2 rounded-md"><TfiReload /></div>
                                            <div className="border-[1px] text-[16px] hover:text-blue-400 border-blue-300 p-2 rounded-md"><PiNotePencilLight /></div>
                                            <div className="border-[1px] text-[16px] hover:text-blue-400 border-blue-300 p-2 rounded-md"><RiDeleteBinLine /></div>
                                        </a>
                                    </div>
                                </div>
                                <div>
                                    <hr className='w-full'/>
                                </div>
                                {selectedScheduleRow == scan.id ? (
                                    scheduletext.map((txt,index)=>(
                                        <div className='flex mt-2 gap-4'>
                                            <div className='flex flex-col items-end'>
                                                <hr className='w-[150px]'/>    
                                                <span className='md:text-xs text-[10px] text-gray-500 mr-3 mt-2'>{txt.name}</span>
                                            </div> 
                                            <div className='flex flex-col items-start'>   
                                                <span className={`md:text-xs text-[10px] font-bold mt-2 ${index === scheduletext.length - 1 ? 'mb-5' : ''}`}>
                                                    {txt.value}
                                                </span>
                                            </div> 
                                        </div> 
                                    ))
                                ):(
                                    <span></span>
                                )}
                            </div>
                        ))}
                    </div>
                </div>
            </div>
        </div>
    </div>
  )
}

export default ScanComppaer1
