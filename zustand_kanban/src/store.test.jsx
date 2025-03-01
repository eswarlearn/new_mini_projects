/*it is to be mounted or custome hooked in some components*/

import { expect, test, vi } from "vitest"
import { useStore } from "./Store";
import { useEffect } from "react";
import { render } from "@testing-library/react";

function TestComponent({selector, effect}){
    const items = useStore(selector);

    useEffect(()=>effect(items),[items]);

    return null;
}

// test("sample",()=>{
//     expect(1).toEqual(1);
// });

test("should return default value ath the startng",()=>{
    const selector =(store) => store.tasks;
    const effect = vi.fn(); //it is called mocked function , it like normal function can be called just for testing 
    render(<TestComponent selector={selector} effect={effect}/>);
    expect(effect).toHaveBeenCalledWith([]);
} )

test("should add an item to store and return the effect",()=>{
    const selector =(store) => ({tasks:store.tasks, addTask:store.addTask});
    const effect = vi.fn(); //it is called mocked function , it like normal function can be called just for testing 
    render(<TestComponent selector={selector} effect={effect}/>);
    expect(effect).toHaveBeenCalledWith([]);
} )