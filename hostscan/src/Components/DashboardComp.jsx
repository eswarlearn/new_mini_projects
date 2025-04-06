import React, { useContext } from 'react'
import { FaPlus } from "react-icons/fa6";
import { DataContext } from './DataContext';

const DashboardComp = () => {
    var {hiddencontent,setHiddenContent}=useContext(DataContext)
  return (
    <div className={`px-6 mt-1 md:mt-12 ${hiddencontent?'md:ml-[200px]':'md:ml-[90px]'} w-full gap-3 flex flex-col`}>
        <div className='flex md:flex-row flex-col gap-3 items-start md:justify-between md:items-center justify-start'>
            <div>
                <h1 className='mt-20 text-5xl font-bold'>Dashboard </h1>
            </div>
            <div className='    flex gap-2 md:hidden block'>
                <button className='bg-black text-white px-6 py-3 items-center justify-center rounded-[4px] flex gap-2 w-auto md:w-max'>
                    <FaPlus className='mt-1'/>
                    <span className='text-[15px] font-bold'>New Scan</span>
                </button>
                <button className='border  items-center px-4 py-3 justify-center rounded-[4px] flex gap-2 w-auto md:w-max'
                    style={{backgroundColor:'whitesmoke'}}>
                    <FaPlus className='mt-1'/>
                    <span className='text-[15px]'>Add Targets</span>
                </button>
            </div>
            <div className='flex items-center gap-3 mt-4 md:mt-24 '>
                <span className='text-sm'>0 scans in progress</span>
                <div className='mx-2 h-6 border-l border-gray-500'></div>
                <span className='text-sm'>3 scheduled scans</span>
            </div>
        </div>
        <div className='flex items-end mt-2 md:mt-6'>
            <span className='font-medium text-2xl'>Risks detected</span>
            <span className='ml-5 text-sm text-gray-400'>Total: 0</span>
        </div>
        <div className='flex flex-col lg:flex-row gap-2 lg:gap-4 '>
            <div className='flex gap-2 lg:gap-4 md:flex-row flex-col w-full'>
                <div className='bg-fuchsia-400 py-4 rounded-lg flex items-center justify-between px-3 w-full lg:relative lg:flex-col lg:items-start'>
                    <h2 className='text-2xl lg:text-lg'>Critical</h2>
                    <h2 className='text-4xl lg:text-6xl lg:absolute lg:bottom-2 lg:right-3'>0</h2>
                </div>
                <div className='bg-red-500 py-4  rounded-lg flex items-center justify-between px-3 w-full lg:relative lg:flex-col lg:items-start'>
                    <h2 className='text-2xl lg:text-lg'>High</h2>
                    <h2 className='text-4xl lg:text-6xl lg:absolute lg:bottom-2 lg:right-3'>0</h2>
                </div>
            </div>
            <div className='flex gap-2 lg:gap-4 md:flex-row flex-col w-full'>
                <div className='bg-orange-300 py-4  rounded-lg flex items-center justify-between px-3 w-full lg:relative lg:flex-col lg:items-start'>
                    <h2 className='text-2xl lg:text-lg'>Medium</h2>
                    <h2 className='text-4xl lg:text-6xl lg:absolute lg:bottom-2 lg:right-3'>0</h2>
                </div>
                <div className='bg-yellow-300 py-4  rounded-lg flex items-center justify-between px-3 w-full lg:relative lg:flex-col lg:items-start'>
                    <h2 className='text-2xl lg:text-lg'>Low</h2>
                    <h2 className='text-4xl lg:text-6xl lg:absolute lg:bottom-2 lg:right-3'>0</h2>
                </div>
            </div>
            <div className={`flex gap-2 lg:gap-4 md:flex-row lg:flex-col flex-col w-full ${hiddencontent?'lg:w-[225px]':'lg:w-[250px]'}`}>
                <div className={`bg-green-400 py-4 rounded-lg flex items-center justify-between px-3 w-full ${hiddencontent?'lg:w-[225px]':'lg:w-[250px]'}`}>
                    <h2 className='text-2xl lg:text-lg'>Accepted</h2>
                    <h2 className='text-4xl lg:text-2xl'>0</h2>
                </div>
                <div className={`bg-slate-200 py-4 rounded-lg flex items-center justify-between px-3 w-full ${hiddencontent?'lg:w-[225px]':'lg:w-[250px]'}`}>
                    <h2 className='text-2xl lg:text-lg'>Closed</h2>
                    <h2 className='text-4xl lg:text-2xl'>0</h2>
                </div>
            </div>
        </div>
    </div>
  )
}
//64
export default DashboardComp
