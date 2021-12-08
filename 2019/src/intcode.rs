use std::collections::HashMap;

pub struct IntCode {
    memory: Memory,
    pc: i64
}

struct Memory {
    data: HashMap<i64, i64>
}

impl IntCode {
    pub fn new(program: Vec<i64>) -> IntCode {
        IntCode {
            memory: Memory::new(program),
            pc: 0
        }
    }

    pub fn get_mem(&self, n: i64) -> &i64 {
        self.memory.get(n)
    }

    pub fn set_mem(&mut self, n: i64, val: i64) {
        self.memory.set(n, val);
    }

    fn advance_pc(&mut self, n: i64) {
        self.pc += n;
    }

    pub fn tick(&mut self) -> bool {
        let instruction = &self.get_mem(self.pc);
        if **instruction == 1 {
            let a1 = &self.get_mem(self.pc + 1);
            let a2 = &self.get_mem(self.pc + 2);
            let a3 = &self.get_mem(self.pc + 3);
            let v1 = **&self.get_mem(**a1);
            let v2 = **&self.get_mem(**a2);
            let result = v1 + v2;
            self.set_mem(**a3, result);
            self.advance_pc(4);
            
        }
        else if **instruction == 2 {
            let a1 = &self.get_mem(self.pc + 1);
            let a2 = &self.get_mem(self.pc + 2);
            let a3 = &self.get_mem(self.pc + 3);
            let v1 = **&self.get_mem(**a1);
            let v2 = **&self.get_mem(**a2);
            &self.set_mem(**a3, v1 * v2);
            self.advance_pc(4);
        }
        else if **instruction == 99 {
            return false;
        }
        true
    }
}

impl Memory {
    fn new(data: Vec<i64>) -> Memory {
        let mut m = HashMap::new();
        for (i,v) in data.into_iter().enumerate() {
            m.insert(i as i64, v as i64);
        }
        Memory {
            data: m
        }
    }

    fn get(&self, n: i64) -> &i64 {
        match self.data.get(&n) {
            Some(val) => val,
            None => &0
        }
    }

    fn set(&mut self, n: i64, val: i64) {
        self.data.insert(n, val);
//        self.data[&n] = val;
//        *self.data.entry(n).or_default() = val;
    }
}