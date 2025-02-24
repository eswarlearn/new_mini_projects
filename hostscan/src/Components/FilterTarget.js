import React, { useState, useContext, useEffect, useRef } from 'react'
import { DataContext } from './DataContext';
import { FaPlus } from "react-icons/fa6";
import { IoFilterSharp } from "react-icons/io5";
import { BsThreeDots } from "react-icons/bs";
import { SlArrowRight } from "react-icons/sl";
import { SlArrowDown } from "react-icons/sl";
import AddTarget from './AddTarget';
import { IoMdClose } from "react-icons/io";
import { MdOutlineKeyboardArrowDown } from "react-icons/md";
import { BiTrash } from "react-icons/bi";
import { FaWifi } from "react-icons/fa";
import { PiNotePencilLight } from "react-icons/pi";
import { BiError } from "react-icons/bi";

const FilterTarget = () => {
  var modalRef = useRef(null)
  var { filteropen, setFilterOpen, targets, setTargets } = useContext(DataContext)
  var [filterID, setFilterID] = useState(null)
  const [filterDropdown, setFilterDropdown] = useState([]);

  const filterTargets = [
    { id: 1, name: 'Label', key: 'label' },
    { id: 2, name: 'Source', key: 'source' },
    { id: 3, name: 'Source Type', key: 'source' },
    { id: 4, name: 'Tags', key: 'tag' },
    { id: 5, name: 'Target', key: 'target_url' },
  ];

  var togglefilter = (e) => {
    setFilterID(filterID == e ? null : e)
  }
  var filterdropdown = [
    // {id:1,name:'fdfhjgd'},
    // {id:2,name:'fdfhjgd'},
  ]

  const toggleFilter = (id, key) => {
    setFilterID(filterID === id ? null : id);

    if (filterID !== id) {
      const uniqueValues = [...new Set(targets.map(target => target[key]).flat())].filter(Boolean);
      setFilterDropdown(uniqueValues);
    } else {
      setFilterDropdown([]);
    }
  };

  const [inputs, setInputs] = useState([{ tag: [] }]);

  const clearTags = (index) => {
    setInputs(prevInputs =>
      prevInputs.map((input, i) => (i === index ? { ...input, tag: [] } : input))
    );
  };

  const handleKeyDown = (index, e) => {
    if ((e.key === "Enter" || e.key === ",") && e.target.value.trim() !== "") {
      e.preventDefault();
      const newTag = e.target.value.trim();

      if (!inputs[index].tag.includes(newTag)) {
        setInputs(prevInputs =>
          prevInputs.map((input, i) =>
            i === index && input.tag.length < 5
              ? { ...input, tag: [...input.tag, newTag] }
              : input
          )
        );
      }

      e.target.value = "";
    }
  };

  const removeTag = (inputIndex, tagIndex) => {
    setInputs(prevInputs =>
      prevInputs.map((input, i) =>
        i === inputIndex
          ? { ...input, tag: input.tag.filter((_, j) => j !== tagIndex) }
          : input
      )
    );
  };

  return (
    <div>
      <div className='flex items-center justify-between'>
        <span className='text-xl'>Filter targets</span>
        <IoMdClose onClick={() => setFilterOpen(true)} className='text-xl' />
      </div>
      <div className='flex'>
        <div className='flex flex-col'>
          {/* <div ref={modalRef} className='mt-5'>
                  {filtertargets.map((valuee,index) => (
                      <div key={index} className='mt-4'>
                        <p className='text-xs bg-white flex items-start ml-3 mt-[-7px] fixed'>{valuee.name}</p>
                        <div onClick={()=>togglefilter(valuee.id)} className='flex items-center mt-2 border border-gray-300 rounded-md px-2 focus-within:border-2 
                          focus-within:border-blue-500'>
                          <input type='text' className='w-full py-2 outline-none' placeholder='(all)'/>
                          <div className='mx-2 h-6 border-l border-gray-300'></div>
                          <MdOutlineKeyboardArrowDown className='text-gray-400 font-bold text-xl' />
                        </div>
                        {filterID == valuee.id ? (
                          <div className={`py-2 border rounded-[7px] mt-2`}>
                            <span className='text-gray-400'>
                              {filterdropdown.length == 0 ? (
                                'No options'
                              ):(
                                filterdropdown.map((item,index) => (
                                  <div key={index} className='flex flex-col items-start p-3'>
                                    <span className=''>{item.name}</span>
                                  </div>
                                ))
                              )}
                            </span>
                          </div>
                        ):(
                          <span></span>
                        )}
                      </div>
                  ))}
                  </div> */}
          <div ref={modalRef} className='mt-5'>
            {filterTargets.map((filter, index) => (
              <div key={index} className='mt-4'>
                <p className='text-xs bg-white flex items-start ml-3 mt-[-7px] fixed'>{filter.name}</p>
                <div onClick={() => toggleFilter(filter.id, filter.key)} className='flex items-center mt-2 border border-gray-300 rounded-md px-2 focus-within:border-2 focus-within:border-blue-500'>
                  <input type='text' className='w-full py-2 outline-none' placeholder='(all)' readOnly />
                  <div className='mx-2 h-6 border-l border-gray-300'></div>
                  <MdOutlineKeyboardArrowDown className='text-gray-400 font-bold text-xl' />
                </div>
                {filterID === filter.id && (
                  <div className='py-2 border rounded-[7px] mt-2'>
                    {filterDropdown.length === 0 ? (
                      <span className='text-gray-400'>No options</span>
                    ) : (
                      filterDropdown.map((item, idx) => (
                        <div key={idx} className='flex flex-col items-start p-3'>
                          <span>{item}</span>
                        </div>
                      ))
                    )}
                  </div>
                )}
              </div>
            ))}

            {inputs.map((input, index) => (
              <div key={index} className='flex flex-wrap items-start w-full sm:w-auto mt-4'>
                <span>Tags</span>
                <div className='relative w-full flex items-center border py-2 px-3 rounded-md md:w-[180px] cursor-text'>
                  <div className='flex flex-wrap w-[90%] gap-1 truncate'>
                    {input.tag.map((tag, tagIndex) => (
                      <div key={tagIndex} className='bg-gray-200 px-2 py-1 rounded flex items-center gap-1'>
                        <span className='truncate max-w-[80px]'>{tag}</span>
                        <button type='button' onClick={() => removeTag(index, tagIndex)} className='text-red-500 font-bold'>✖</button>
                      </div>
                    ))}
                    {input.tag.length < 5 && (
                      <input
                        type='text'
                        placeholder='Up to 5 tags'
                        className='outline-none flex-grow min-w-[80px]'
                        onKeyDown={(e) => handleKeyDown(index, e)}
                      />
                    )}
                  </div>
                  {input.tag.length > 0 && (
                    <button type='button' onClick={() => clearTags(index)} className='absolute right-2 text-red-600 font-bold'>✖</button>
                  )}
                </div>
              </div>
            ))}
          </div>
          <hr className='mt-6' />
          <div className='mt-6 flex gap-6 font-medium'>
            <button onClick={() => setFilterOpen(true)} className='border items-center bg-stone-100 flex px-6 rounded-md gap-3 py-3'>
              <BiTrash className='' />
              <span className='text-sm'>Clear Filters</span>
            </button>
            <button className='bg-black text-sm py-3 px-12 text-white hover:bg-stone-900'>Apply</button>
          </div>
        </div>
      </div>
    </div>
  )
}

export default FilterTarget
