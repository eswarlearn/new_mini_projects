import React, { useContext } from 'react'
import DashHeader from '../Components/DashHeader'
import DashNav from '../Components/DashNav'
import Riskcomp from '../Components/Riskcomp'
import { DataContext } from '../Components/DataContext'

const Risks = () => {
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
        <Riskcomp/>
      </div>
    </div>
  )
}

export default Risks
