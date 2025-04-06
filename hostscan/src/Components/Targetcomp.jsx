import React, { useState, useContext, useEffect, useRef } from 'react'
import { DataContext } from './DataContext';
import { FaPlus } from "react-icons/fa6";
import { IoFilterSharp } from "react-icons/io5";
import { BsThreeDots } from "react-icons/bs";
import { SlArrowRight } from "react-icons/sl";
import { SlArrowDown } from "react-icons/sl";
import AddTarget from './AddTarget';
import FilterTarget from './FilterTarget';
import { BiTrash } from "react-icons/bi";
import { IoMdClose } from "react-icons/io";
import { FaWifi } from "react-icons/fa";
import { PiNotePencilLight } from "react-icons/pi";
import { BiError } from "react-icons/bi";
import dayjs from 'dayjs';
import utc from 'dayjs/plugin/utc';
import timezone from 'dayjs/plugin/timezone';
import relativeTime from 'dayjs/plugin/relativeTime';
import { AiFillCaretUp, AiFillCaretDown } from "react-icons/ai";
import axios from "axios";

dayjs.extend(utc);
dayjs.extend(timezone);
dayjs.extend(relativeTime);

const Targetcomp = () => {
  var modalRef = useRef(null)
  var ellipRef = useRef(null)
  var inputRef = useRef(null);
  var btnRef = useRef({});
  var deleteRef = useRef({});
  var editRef = useRef({});
  var { hiddencontent, setHiddenContent, targets, setTargets } = useContext(DataContext);
  var [scanDetails, setScanDetails] = useState(true);
  var [selectedScheduleRow, setSelectedScheduleRow] = useState(null);
  var { addTargets, setAddTargets } = useContext(DataContext);
  var { filteropen, setFilterOpen } = useContext(DataContext);
  var [deleteTarget, setDeleteTarget] = useState(true);
  var [editTarget, setEditTarget] = useState(true);
  var [selectEllip, setSelecteEllip] = useState(null);
  const [editScan, setEditScan] = useState([{ target_url: "", tag: [], label: "" }]);
  var [indexvalues, setIndexValues] = useState(null);

  const user_id = localStorage.getItem("user_id");
  const Authorization = axios.defaults.headers.common["Authorization"];


  const fetchTargets = async () => {
    try {
      const response = await axios.post(
        'http://localhost:5000/api/view/targets', { user_id: user_id }, 
        {
          headers: {
            'Content-Type': 'application/json',
            'Authorization': Authorization,
          },
          withCredentials: true,
        }
      ); 
      
      const formattedData = response.data.map((data, index) => ({
        id: index,
        user_id: data.user_id,
        target_url: data.target_url,
        target_id: data.target_id,
        tag: typeof data.tag === 'string' ? JSON.parse(data.tag) : data.tag,
        label: data.label,
        source: data.source,
        frequency: data.frequency || '',
        days: data.target_created_date || null,
      }));

      setTargets(formattedData);
    } catch (err) {
      console.error('Error fetching data:', err);
    }
  };

  useEffect(() => {
    fetchTargets();
  }, []);

  var toggleellip = (e) => {
    setIndexValues(e)
    setSelecteEllip(e == selectEllip ? null : e)
  }


  var targetData = [
    { id: 1, name: "Manual", target: "", frequency: "2", days: "23 days ago" },
  ];
  var scheduletext = [
    { id: 1, name: 'schedule', value: 'MONTHLY' },
    { id: 2, name: 'id', value: '6790bb54ecc390814192d35c' },
    { id: 3, name: 'scan_type', value: 'OPENVAS' },
    { id: 4, name: 'email_notification', value: 'ALWAYS' },
    { id: 5, name: 'next_run', value: '2025-02-22T09:33:08.572Z' },
    { id: 6, name: 'created_at', value: '2025-01-22T09:33:08.573Z' },
  ]
  var togglescheduletarget = (e) => {
    setSelectedScheduleRow(e == selectedScheduleRow ? null : e)
    setScanDetails(!scanDetails)
  }
  useEffect(() => {
    const handleClickOutside = (event) => {
      const isClickInsideModal = modalRef.current && modalRef.current.contains(event.target);
      const isClickInsideEllip = ellipRef.current && ellipRef.current.contains(event.target);
      var clickdelete = Object.values(deleteRef.current).some(ref => ref && ref.contains(event.target));
      const isClickInsideAnyButton = Object.values(btnRef.current).some(ref => ref && ref.contains(event.target));
      const isClickeditButton = Object.values(editRef.current).some(ref => ref && ref.contains(event.target));

      if (!isClickInsideModal && !isClickInsideEllip && !isClickInsideAnyButton && !clickdelete && !isClickeditButton) {
        setFilterOpen(true);
        setSelecteEllip(null);
        setDeleteTarget(true)
        setEditTarget(true)
      }
    };
    document.addEventListener("mousedown", handleClickOutside);
    return () => {
      document.removeEventListener("mousedown", handleClickOutside);
    };
  }, []);

  const clickEditBtn = (e) => {
    setEditTarget(false);
    const targetData = targets[e];
    setEditScan({
      id: targetData?.id,
      user_id: targetData?.user_id || '',
      target_url: targetData?.target_url || '',
      target_id: targetData?.target_id || '',
      tag: targetData?.tag || [],
      label: targetData?.label || '',
      days: targetData?.days,
    });
  };

  const handleChange = (e) => {
    const { name, value } = e.target;
    setEditScan((prev) => ({ ...prev, [name]: value }));
  };


  var clickdeletebtn = (e) => {
    setDeleteTarget(!deleteTarget)
  }
  var clickscanbtn = (e) => {

  }
  var cancalDeleteChange = () => {
    setDeleteTarget(!deleteTarget)
    setSelecteEllip(null)
  }
  var handleEditChange = () => {
    setEditTarget(!editTarget)
    setSelecteEllip(null)
  }

  const handleKeyDown = (e) => {
    if ((e.key === 'Enter' || e.key === ',') && editScan.inputValue?.trim() !== '') {
      e.preventDefault();
      if (editScan.tag.length < 5) {
        setEditScan((prev) => ({
          ...prev,
          tag: [...prev.tag, prev.inputValue.trim()],
          inputValue: ''
        }));
      }
    }
  };

  const handleBlur = () => {
    console.log('in blur handle');
    
    if (editScan.inputValue?.trim() !== '') {
      setEditScan((prev) => ({
        ...prev,
        tag: [...prev.tag, prev.inputValue.trim()], // Add the input as a tag
        inputValue: '' // Clear the input field
      }));
    }
  };
  

  const removeTag = (index) => {
    setEditScan((prev) => ({
      ...prev,
      tag: prev.tag.filter((_, i) => i !== index)
    }));
  };

  const clearTags = () => {
    setEditScan((prev) => ({ ...prev, tag: [] }));
  };
  // Assuming you're using Axios for API requests

  const handleEditSubmit = async (updatedData) => {
    try {
      const response = await axios.post(
        'http://localhost:5000/api/targets/edit',
        updatedData, // The data to send in the request body
        {
          headers: {
            'Content-Type': 'application/json',
            'Authorization': Authorization,
          },
          withCredentials: true, // This ensures credentials are sent with the request
        }
      );

      if (response.status === 200) {
        const updatedResult = targets.map((item) =>
          item.id === updatedData.id ? { ...item, ...updatedData } : item
        );
        setTargets(updatedResult);
        setEditTarget(true);
        console.log('Target updated successfully!');
      } else {
        const errorMessage = response.data.message || 'Error updating target';
        console.error(errorMessage);
        alert(errorMessage);
      }
    } catch (error) {
      console.error('Error:', error);
      alert(error.response?.data?.message || 'Server Error (500)');
    }
  };

  // Handle Delete with Error Handling (POST)
  const handleDeleteChange = async (target) => {
    try {
      const response = await axios.post(`http://localhost:5000/api/targets/delete`, { target_id: target.target_id },
        {
          headers: {
            'Content-Type': 'application/json',
            'Authorization': Authorization,
          },
          withCredentials: true, // This ensures credentials are sent with the request
        });

      if (response.status === 200) {
        const updatedTargets = targets.filter((item) => item.id !== target.id);
        setTargets(updatedTargets);
        setDeleteTarget((prev) => !prev);
        console.log('Target deleted successfully!');
      } else {
        const errorMessage = response.data.message || 'Error deleting target';
        console.error(errorMessage);
        alert(errorMessage);
      }
    } catch (error) {
      console.error('Error:', error);
      alert(error.response?.data?.message || 'Server Error (500)');
    }
  };




  // *******************************


  // // Sort Function
  // const sortedTargets = [...targetResult].sort((a, b) => {
  //   if (sortConfig.key) {
  //     const valueA = a[sortConfig.key] || '';
  //     const valueB = b[sortConfig.key] || '';

  //     if (typeof valueA === 'string') {
  //       return sortConfig.direction === 'asc'
  //         ? valueA.localeCompare(valueB)
  //         : valueB.localeCompare(valueA);
  //     } else {
  //       return sortConfig.direction === 'asc' ? valueA - valueB : valueB - valueA;
  //     }
  //   }
  //   return 0;
  // });


  // // Toggle Sorting
  // const handleSort = (key) => {
  //   let direction = 'asc';
  //   if (sortConfig.key === key && sortConfig.direction === 'asc') {
  //     direction = 'desc';
  //   }
  //   setSortConfig({ key, direction });
  // };



  return (
    <div className={`mt-20 px-4 flex ${hiddencontent ? 'md:ml-[200px]' : 'md:ml-[90px]'} flex-col gap-8 md:mt-32 w-full`}>
      <div className={`absolute top-0 left-0 w-full h-full ${addTargets ? 'hidden' : 'block bg-gray-500 bg-opacity-50 z-50 overflow-auto'}`}>
        <div className="flex justify-center items-center min-h-screen">
          <AddTarget />
        </div>
      </div>
      <div ref={modalRef} className={`fixed ${filteropen ? 'hidden' : 'block'} max-sm:w-full mt-[-25px] right-0 sm:right-4 border px-5 py-4 rounded-[7px] bg-white flex flex-col `}>
        <FilterTarget />
      </div>
      <div className={``}>
        <div className='gap-2 flex md:flex-row flex-col items-start justify-start md:items-center md:justify-between'>
          <div className=''>
            <h1 className='font-bold text-4xl'>Targets</h1>
          </div>
          <div className='flex gap-2'>
            <button onClick={() => setAddTargets(!addTargets)} className='bg-black block text-white px-5 py-3 items-center justify-center rounded-[4px] flex gap-2 w-auto md:w-max'>
              <FaPlus className='mt-1' />
              <span className='text-[15px] font-bold'>Add Targets</span>
            </button>
            <button onClick={() => setFilterOpen(!filteropen)} className='border items-center px-4 py-3 justify-center rounded-[4px] flex gap-2 w-auto md:w-max'
              style={{ backgroundColor: 'whitesmoke' }}>
              <IoFilterSharp className='mt-1' />
              <span className='text-[15px]'>Filters</span>
            </button>
          </div>
        </div>
        <div className='border-2 hidden border-gray-300 rounded-lg p-8 border-dashed'>
          <span className='text-gray-500 text-sm'>No data to display</span>
        </div>
        <div className="mt-12 flex flex-col">
          <div className="space-y-1">
            {/* First Row (Header) */}
            <div className="flex w-full items-center text-xs font-medium justify-between">
              <div className="w-[180px] lg:w-[300px] flex-shrink-0 flex items-center"
              // onClick={() => handleSort('target_url')}
              >
                <a className="ml-4">Targets</a>
                {/* {sortConfig.key === 'target_url' && (sortConfig.direction === 'asc' ? <AiFillCaretUp /> : <AiFillCaretDown />)} */}

              </div>
              <div className="grid grid-cols-3 md:grid-cols-4 flex-1 text-end">
                <div className="cursor-pointer"
                // onClick={() => handleSort('source')}
                >
                  <a>Source/Type</a>
                  {/* {sortConfig.key === 'source' && (sortConfig.direction === 'asc' ? <AiFillCaretUp /> : <AiFillCaretDown />)} */}
                </div>

                <div className="cursor-pointer"
                // onClick={() => handleSort('frequency')}
                >
                  <a>Last Scanned</a>
                  {/* {sortConfig.key === 'frequency' && (sortConfig.direction === 'asc' ? <AiFillCaretUp /> : <AiFillCaretDown />)} */}
                </div>


                <div className="hidden md:block cursor-pointer"
                // onClick={() => handleSort('days')}
                >
                  <a className="hidden md:block"> Created </a>
                  {/* {sortConfig.key === 'days' && (sortConfig.direction === 'asc' ? <AiFillCaretUp /> : <AiFillCaretDown />)} */}
                </div>

              </div>
            </div>
            {/* Horizontal Line */}
            <hr className="w-full mt-2" />

            {/* shorting row 
            {sortedTargets && Array.isArray(sortedTargets) && sortedTargets.map((scan, index) => (
              <div key={scan.id} className="flex w-full hover:bg-gray-100 items-center text-xs font-medium justify-between">
                <div className="w-[180px] lg:w-[300px] ml-4 flex-shrink-0 text-left">
                  <div className="font-mono text-sm leading-5 text-slate-400">{scan.target_url}</div>
                  <div className="font-medium">{scan.label}</div>
                </div>
                <div className="grid grid-cols-3 md:grid-cols-4 flex-1 items-center text-end">
                  <a>{scan.source}</a>
                  <a>{scan.frequency}</a>
                  <a className="hidden md:block">{scan.days ? dayjs.tz(scan.days, "America/New_York").fromNow() : "N/A"}</a>
                </div>
              </div>
            ))}*/}







            {/* Data Rows */}
            <div className="space-y-1">
              {targets && Array.isArray(targets) && targets.map((scan, index) => (
                <div key={scan.id} className="flex flex-col">
                  <div className={`fixed top-0 left-0 w-full h-full ${deleteTarget ? 'hidden' : 'block bg-gray-500 bg-opacity-50 z-50 overflow-auto'} rounded-lg flex justify-center items-center`}>
                    <div ref={(el) => (deleteRef.current[scan.id] = el)} className="flex flex-col items-center w-[550px] bg-white overflow-auto p-6 rounded-lg shadow-lg ">
                      <div className="flex justify-between w-[850px] max-w-full mx-auto items-center mt-3">
                        <div className='flex items-center'>
                          <div className='rounded-[100%] bg-rose-200 p-2'><BiError className='text-xl text-red-500' /></div>
                          <h2 className="text-lg ml-2">Delete Target</h2>
                        </div>
                        <button onClick={cancalDeleteChange}><IoMdClose className="text-gray-400 text-2xl cursor-pointer" /></button>
                      </div>
                      <div className="mt-3 w-[850px] max-w-full sm:text-left px-4">
                        <p className="font-medium text-sm text-center sm:text-left">
                          IPv4 / DNS Name: <span className='font-black'>{targets?.[indexvalues]?.target_url}</span>
                        </p>
                        <p className='mt-2 text-sm font-medium'>Are you sure you want to delete this target?The risks for this target will also be removed.</p>
                        <p className='mt-2 text-sm font-medium'>
                          Note that if the target was already scanned this month then it will count towards your target capacity until the end of the period.View
                          your <span className='underline'>account usage dashboard</span> for more information.
                        </p>
                      </div>
                      <div className="mt-5 flex justify-between w-[850px] max-w-full mx-auto">
                        <button onClick={() => { handleDeleteChange(targets?.[indexvalues]) }} className="bg-red-600 py-2 rounded-md border border-red-700 text-white font-bold px-20">
                          Delete
                        </button>
                        <button onClick={cancalDeleteChange} className="bg-white py-2 rounded-md border border-black font-bold px-20">
                          Cancel
                        </button>
                      </div>
                    </div>
                  </div>
                  <div className={`fixed top-0 left-0 w-full h-full ${editTarget ? 'hidden' : 'block bg-gray-500 bg-opacity-50 z-50 overflow-auto'} rounded-lg flex justify-center items-center`}>
                    <div ref={(el) => (editRef.current[scan.id] = el)} className="flex flex-col items-center w-[800px] bg-white overflow-auto p-3 rounded-lg shadow-lg ">
                      <div className="flex justify-between w-[850px] max-w-full mx-auto items-center mt-3">
                        <h2 className="text-2xl font-bold">Edit Target</h2>
                        <button onClick={handleEditChange}><IoMdClose className="text-gray-400 text-2xl cursor-pointer" /></button>
                      </div>
                      <div className="flex flex-col items-center justify-center sm:flex-row gap-4 w-[850px] max-w-full mt-3">
                        <div className="flex items-start flex-col w-full sm:w-auto">
                          <span>Target *</span>
                          <input type="text" name="target_url" value={editScan.target_url} placeholder="IPv4, Domain, URL or Public CIDR" className="border py-2 md:w-[260px] rounded-md px-3 w-full" onChange={handleChange} />
                        </div>
                        <div className="flex items-start flex-col w-full sm:w-auto">
                          <span>Tags</span>
                          <div className="relative w-full flex items-center border py-2 px-3 rounded-md md:w-[200px] cursor-text"
                            onClick={() => inputRef.current?.focus()}
                          >
                            {/* Display Tags */}
                            <div className="flex flex-wrap w-[90%] gap-1">
                              {editScan.tag && Array.isArray(editScan.tag) && editScan.tag.map((tag, index) => (
                                <div key={index} className="bg-gray-200 px-2 py-1 rounded flex items-center gap-1">
                                  <span className="truncate max-w-[80px]">{tag}</span>
                                  <button type="button" onClick={() => removeTag(index)} className="text-red-500 font-bold">
                                    ✖
                                  </button>
                                </div>
                              ))}

                              {/* Input Field for New Tags */}
                              {editScan.tag && Array.isArray(editScan.tag) && editScan.tag.length < 5 && (
                                <input
                                  ref={inputRef}
                                  type="text"
                                  placeholder="Add tags..."
                                  name="inputValue"
                                  value={editScan.inputValue || ''}
                                  className="outline-none flex-grow min-w-[80px] bg-transparent"
                                  onChange={handleChange}
                                  onKeyDown={handleKeyDown}
                                  onBlur={handleBlur}
                                />
                              )}
                            </div>

                            {/* Clear All Tags Button */}
                            {editScan.tag && Array.isArray(editScan.tag) && editScan.tag.length > 0 && (
                              <button type="button" onClick={clearTags} className="absolute right-2 text-red-600 font-bold">
                                ✖
                              </button>
                            )}
                          </div>


                        </div>
                        <div className="flex items-start flex-col w-full sm:w-auto">
                          <span>Label</span>
                          <input type="text" name="label" value={editScan.label} placeholder="e.g. my-example-server" className="border py-2 rounded-md px-3 w-full md:w-[260px]" onChange={handleChange} />
                        </div>
                      </div>
                      <div className="mt-5 flex justify-end w-[850px] max-w-full mx-auto">
                        {/* <button btnIndex ={index}  onClick={handleEditSubmit} className="bg-blue-500 py-3 rounded-md border border-blue-700 text-white font-bold px-3">
                          Submit
                        </button> */}
                        <button
                          onClick={() =>
                            handleEditSubmit({
                              id: editScan.id,
                              target_url: editScan.target_url,
                              tag: editScan.tag,
                              label: editScan.label,
                              days: editScan?.days,
                              target_id: editScan?.target_id || '',

                            })
                          }
                          className="bg-blue-500 py-3 rounded-md border border-blue-700 text-white font-bold px-3"
                        >
                          Submit
                        </button>

                      </div>
                    </div>
                  </div>
                  {/* Data Row */}
                  <div className="flex w-full hover:bg-gray-100 items-center text-xs font-medium justify-between">
                    <div className="w-[180px] lg:w-[300px] ml-4 flex-shrink-0 text-left">
                      <div className="font-mono text-sm leading-5 text-slate-400">{scan.target_url}</div>
                      <div className="font-medium">{scan.label}</div>
                      {scan.tag && Array.isArray(scan.tag) && scan.tag.map((tag, inxTg) => (
                        <span key={inxTg} className="bg-gray-200 mr-1 rounded-full border border-gray-400 px-1 gap-2 items-center py-0.5 text-black font-medium text-xxs tracking-[2px] uppercase max-w-64 truncate !text-nowrap">{tag}</span>
                      ))}
                    </div>
                    <div className="grid grid-cols-3 md:grid-cols-4 flex-1 items-center text-end">
                      <a>{scan.source}</a>
                      <a>{scan.frequency}</a>
                      <a className="hidden md:block">
                        {/* {scan.days} */}
                        {(scan.days ? dayjs.tz(scan.days, "America/New_York").fromNow() : "N/A")}
                      </a>
                      <a className="flex items-center justify-end gap-3 ml-20">
                        <div ref={(el) => (btnRef.current[scan.id] = el)} onClick={() => { toggleellip(scan.id) }} className="border-[1px] hover:text-blue-400 text-[16px] p-2 rounded-md">
                          <BsThreeDots />
                        </div>
                      </a>
                      {scan.id == selectEllip ? (
                        <div ref={ellipRef} className={`absolute right-2 rounded-[8px] p-3 mt-44 bg-white border text-md flex gap-2 flex-col`}>
                          <button onClick={() => clickscanbtn(index)} className='flex gap-4 items-center hover:bg-cyan-200 px-2 py-2 font-bold'>
                            <FaWifi />
                            Scan Target
                          </button>
                          <button onClick={() => clickEditBtn(index)} className='flex gap-4 items-center hover:bg-cyan-200 px-2 py-2 font-bold'>
                            <PiNotePencilLight />
                            Edit Target
                          </button>
                          <button onClick={() => clickdeletebtn(index)} className='flex gap-4 items-center hover:bg-rose-200 px-2 py-2 font-bold'>
                            <BiTrash />
                            Delete Target
                          </button>
                        </div>
                      ) : (
                        <span></span>
                      )}
                    </div>
                  </div>
                  <hr className="w-full mt-1" />
                </div>
              ))}
            </div>
          </div>
        </div>
        <div className='flex items-start mt-7'>
          <h1 className='font-bold text-2xl'>Sources</h1>
        </div>
        <div className="flex flex-col">
          <div className="space-y-1">
            {/* First Row (Header) */}
            <div className="flex w-full items-center text-xs font-medium justify-between">
              <div className="w-[180px] lg:w-[300px] flex-shrink-0 flex items-center">
                <a className="ml-20">Type</a>
              </div>
              <div className="grid grid-cols-4 flex-1 text-start">
                <a>Account</a>
                <a className="">Targets</a>
                <a className="">Last Refreshed</a>
              </div>
            </div>
            {/* Horizontal Line */}
            <hr className="w-full mt-2" />
            {/* Data Rows */}
            <div className="space-y-1">
              {targetData.map((scan, index) => (
                <div key={scan.id} className="flex flex-col">
                  {/* Data Row */}
                  <div onClick={() => togglescheduletarget(scan.id)} className="flex w-full hover:bg-gray-100 items-center text-xs font-medium justify-between">
                    <div className="w-[180px] lg:w-[300px] flex-shrink-0 flex items-center">
                      <div className="hover:bg-gray-300 ml-7 px-[18px] py-4 flex items-center justify-center">
                        <button>
                          {selectedScheduleRow == scan.id ? (
                            <SlArrowDown />
                          ) : (
                            <div>
                              <SlArrowRight />
                            </div>
                          )}
                        </button>
                      </div>
                      <span className="">{scan.name}</span>
                    </div>
                    <div className="grid grid-cols-4 flex-1 items-center text-start">
                      <a>{scan.target}</a>
                      <a className=''>{scan.frequency}</a>
                      <a className="">{scan.days}</a>
                    </div>
                  </div>
                  {/* Horizontal Line Between Rows */}
                  <hr className="w-full mt-1" />
                  {selectedScheduleRow == scan.id ? (
                    scheduletext.map((txt, index) => (
                      <div className='flex mt-2 gap-4'>
                        <div className='flex flex-col items-end'>
                          <hr className='w-[150px]' />
                          <span className='md:text-xs text-[10px] text-gray-500 mr-3 mt-2'>{txt.name}</span>
                        </div>
                        <div className='flex flex-col items-start'>
                          <span className={`md:text-xs text-[10px] font-bold mt-2 ${index === scheduletext.length - 1 ? 'mb-5' : ''}`}>
                            {txt.value}
                          </span>
                        </div>
                      </div>
                    ))
                  ) : (
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

export default Targetcomp
