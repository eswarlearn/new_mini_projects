import React, { useContext } from 'react'
import DashHeader from '../Components/DashHeader'
import ScanComppaer1 from '../Components/ScanComppaer1'
import DashNav from '../Components/DashNav'
import { DataContext } from '../Components/DataContext'

const Scan = () => {
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
        <ScanComppaer1/>
      </div>
    </div>
  )
}

export default Scan
