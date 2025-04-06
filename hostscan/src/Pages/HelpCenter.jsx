import React, { useContext } from 'react'
import DashHeader from '../Components/DashHeader'
import DashNav from '../Components/DashNav'
import Helpcentercomp from '../Components/Helpcentercomp'
import { DataContext } from '../Components/DataContext'

const HelpCenter = () => {
  var {hiddencontent,setHiddenContent}=useContext(DataContext)
  return (
    <div className='flex flex-col'>
      <div>
        <DashHeader/>
      </div>
      <div className='flex'>
        <div className={`md:fixed ${hiddencontent ? 'max-md:max-w-full' : 'max-md:max-w-[0px]'}`}>
          <DashNav/>
        </div>
        <Helpcentercomp/>
      </div>
    </div>
  )
}

export default HelpCenter
