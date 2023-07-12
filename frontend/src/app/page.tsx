import EmployeeTable from "@/components/EmployeeTable";

export default function Home() {
  return (
    <main className="flex min-h-screen flex-col items-center p-12 md:p-24">
      <EmployeeTable />
    </main>
  );
}
