import React, { useEffect, useState } from 'react';
import dayjs from 'dayjs';
import timezone from 'dayjs/plugin/timezone';
import utc from 'dayjs/plugin/utc';

dayjs.extend(utc);
dayjs.extend(timezone);

const DateTimeInput = () => {
  const [selectedDate, setSelectedDate] = useState(dayjs().tz('America/New_York'));
  const [hours, setHours] = useState(selectedDate.hour());
  const [minutes, setMinutes] = useState(selectedDate.minute());

  useEffect(() => {
    const now = dayjs().tz('America/New_York');
    setSelectedDate(now); // Ensure the current date is selected properly
    setHours(now.hour()); // Initialize hours correctly
    setMinutes(now.minute()); // Initialize minutes correctly
  }, []);

  const handleDateClick = (day) => {
    setSelectedDate((prev) => prev.date(day)); // Update only the day in selectedDate
  };

  const handleMonthChange = (direction) => {
    const newDate = selectedDate.add(direction, 'month');
    const now = dayjs();
    const oneYearLater = now.add(7, 'month');

    if (newDate.isAfter(oneYearLater) || newDate.isBefore(now.startOf('month'))) return;
    setSelectedDate(newDate);
  };

  const handleTimeChange = (newHours, newMinutes) => {
    setSelectedDate((prev) => prev.hour(newHours).minute(newMinutes));
  };

  const getAvailableHours = () => {
    const now = dayjs().tz('America/New_York');

    console.log('now _getAvailableHours --> ',now);
    return selectedDate.isSame(now, 'day')
      ? Array.from({ length: 24 - now.hour() }, (_, i) => now.hour() + i)
      : Array.from({ length: 24 }, (_, i) => i);
  };

  const getAvailableMinutes = () => {
    const now = dayjs().tz('America/New_York');;
    return selectedDate.isSame(now, 'day') && hours === now.hour()
      ? Array.from({ length: 60 - now.minute() }, (_, i) => now.minute() + i)
      : Array.from({ length: 60 }, (_, i) => i);
  };

  return (
    <div className="flex flex-col items-center justify-center p-4 bg-gray-100 rounded-lg shadow-md">
      {/* ✅ This will now update properly when hours/minutes change */}
      <div className="text-lg font-semibold mb-4">
        {selectedDate.format('DD/MM/YYYY, HH:mm')}
      </div>
      
      <div className="flex items-center mb-4">
        <button onClick={() => handleMonthChange(-1)} className="p-2 bg-gray-300 rounded-l">{'<'}</button>
        <span className="mx-4 font-medium">{selectedDate.format('MMMM YYYY')}</span>
        <button onClick={() => handleMonthChange(1)} className="p-2 bg-gray-300 rounded-r">{'>'}</button>
      </div>
      
      <div className="grid grid-cols-7 gap-2 mb-4">
        {["SU", "MO", "TU", "WE", "TH", "FR", "SA"].map((day) => (
          <div key={day} className="text-center font-medium text-gray-700">{day}</div>
        ))}
        {Array.from({ length: selectedDate.startOf('month').day() }).map((_, i) => (
          <div key={`empty-${i}`} className="text-center"></div>
        ))}
        {Array.from({ length: selectedDate.daysInMonth() }).map((_, i) => {
          const day = i + 1;
          const date = selectedDate.date(day);
          const isDisabled = date.isBefore(dayjs()) || date.isAfter(dayjs().add(7, 'month'));
          return (
            <div
              key={`day-${day}`}
              className={`text-center p-2 rounded cursor-pointer ${
                isDisabled ? 'text-gray-400 cursor-not-allowed' : selectedDate.date() === day ? 'bg-blue-500 text-white' : 'hover:bg-blue-300'
              }`}
              onClick={() => !isDisabled && handleDateClick(day)}
            >
              {day}
            </div>
          );
        })}
      </div>

      <div className="flex space-x-4 mb-4">
        <div className="flex flex-col items-center">
          <label htmlFor="hours" className="text-sm font-medium text-gray-700">Hours</label>
          <select
            id="hours"
            className="w-16 p-2 border rounded"
            value={hours}
            onChange={(e) => {
              const newHours = Number(e.target.value);
              setHours(newHours);
              handleTimeChange(newHours, minutes);
            }}
          >
            {getAvailableHours().map((h) => (
              <option key={h} value={h}>{h}</option>
            ))}
          </select>
        </div>
        <div className="flex flex-col items-center">
          <label htmlFor="minutes" className="text-sm font-medium text-gray-700">Minutes</label>
          <select
            id="minutes"
            className="w-16 p-2 border rounded"
            value={minutes}
            onChange={(e) => {
              const newMinutes = Number(e.target.value);
              setMinutes(newMinutes);
              handleTimeChange(hours, newMinutes);
            }}
          >
            {getAvailableMinutes().map((m) => (
              <option key={m} value={m}>{m}</option>
            ))}
          </select>
        </div>
      </div>
    </div>
  );
};

export default DateTimeInput;
