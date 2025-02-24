import React, { useContext } from 'react'
import DashHeader from '../Components/DashHeader'
import DashNav from '../Components/DashNav'
import Targetcomp from '../Components/Targetcomp'
import { DataContext } from '../Components/DataContext'

const Target = () => {
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
        <Targetcomp/>
      </div>
    </div>
  )
}

export default Target
