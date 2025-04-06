import React, { useContext } from 'react'
import DashHeader from '../Components/DashHeader'
import DashboardComp from '../Components/DashboardComp'
import DashNav from '../Components/DashNav'
import { DataContext } from '../Components/DataContext'
// ${hiddencontent ? 'max-md:max-w-full max-md:fixed' : 'max-md:-translate-x-40'}
const Dashboard = () => {
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
        <DashboardComp/>
      </div>
    </div>
  )
}

export default Dashboard
