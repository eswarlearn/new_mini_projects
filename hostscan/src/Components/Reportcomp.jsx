import React, { useContext,useState } from 'react'
import { DataContext } from './DataContext';

const Reportcomp = () => {
  var {hiddencontent,setHiddenContent}=useContext(DataContext)
  
  return (
    <div className={`mt-20 px-4 ${hiddencontent?'md:ml-[200px]':'md:ml-[90px]'} flex flex-col gap-8 md:mt-32 w-full`}>
      <div className='gap-2 flex md:flex-row flex-col items-start justify-start md:items-center md:justify-between'>
        <div className=''>
          <h1 className='font-bold text-4xl'>Reports</h1>
        </div>
      </div>
    </div>
  )
}

export default Reportcomp
