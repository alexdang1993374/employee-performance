"use client";

import React, { useEffect, useState } from "react";
import axios from "axios";

import { IEmployee, IEmployeeResponse } from "@/types";

const EmployeeTable = () => {
  const [employees, setEmployees] = useState<IEmployee[]>([]);
  console.log(
    "🚀 ~ file: EmployeeTable.tsx:10 ~ EmployeeTable ~ employees:",
    employees
  );
  const [loading, setLoading] = useState<boolean>(true);

  useEffect(() => {
    const fetchData = async () => {
      try {
        const result: IEmployeeResponse = await axios.get(
          "http://localhost:5001/employees"
        );

        setEmployees(result.data.data.data);
        setLoading(false);
      } catch (error) {
        console.error("Error fetching data: ", error);
        setLoading(false);
      }
    };

    fetchData();
  }, []);

  if (loading) {
    return <div>Loading...</div>;
  }

  return <div>Employee table</div>;
};

export default EmployeeTable;
