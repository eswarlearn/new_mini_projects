import React,{createContext,useState} from 'react'

const DataContext=createContext()

const Passingdatas = ({children}) => {
  var [inn,setInn]=useState('in')
  var [hiddencontent,setHiddenContent]=useState(false)
  var [bgnum,setbgnum]=useState(0)
  var [addTargets, setAddTargets]=useState(true)
  const [targets, setTargets] = useState([]); // Store the submitted data from AddTarget.js
  var [emailId, setEmailId] = useState();
  var [userId, setUserId] = useState()
  var [filteropen, setFilterOpen]=useState(true);
  return (
    <DataContext.Provider value={{inn,setInn,hiddencontent,setHiddenContent,bgnum,setbgnum, addTargets,setAddTargets,targets, setTargets,filteropen,setFilterOpen,emailId, setEmailId,userId, setUserId}}>
      {children}
    </DataContext.Provider>
  )
}

export {DataContext, Passingdatas}
