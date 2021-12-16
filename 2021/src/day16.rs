use hex::FromHex;

pub fn part1(input: String) {
    let mut streamer = DataStreamer::new(input);
    let packet = Packet::read(&mut streamer);
    println!("Part 1: {}", total_versions(&packet));
}

pub fn part2(input: String) {
    let mut streamer = DataStreamer::new(input);
    let packet = Packet::read(&mut streamer);
    println!("Part 2: {}", packet.value);
}

fn total_versions(packet: &Packet) ->u64 {
    let mut result = packet.version;
    for child in &packet.children {
        result += total_versions(child);
    }
    result
}

struct Packet {
    version: u64,
    value: u64,
    children: Vec<Packet>,
    binary_length: usize
}

impl Packet {
    fn read(streamer: &mut DataStreamer) -> Self {
        let version = streamer.take(3);
        let type_id = streamer.take(3);
        let value;
        let mut children = vec![];
        let mut binary_length = 6; // Length of header

        if type_id != 4 {
            let length_type_id = streamer.get_bit();
            match length_type_id {
                0 => {
                    let subpacket_bit_length = streamer.take(15) as usize;
                    binary_length += 16 + subpacket_bit_length; // The bit length value, type id and subpackets
                    let mut consumed_bits = 0;
                    while consumed_bits < subpacket_bit_length {
                        let subpacket = Packet::read(streamer);
                        consumed_bits += subpacket.binary_length;
                        children.push(subpacket);
                    }
                },
                1 => {
                    binary_length += 12; // The num subpackets value and the type id
                    let num_subpackets = streamer.take(11);
                    for _ in 0..num_subpackets {
                        let subpacket = Packet::read(streamer);
                        binary_length += subpacket.binary_length;
                        children.push(subpacket);
                    }
                },
                _ => { panic!("Unexpected length_type_id"); }
            }
        }

        match type_id {
            0 => { // Sum
                value = children.iter().map(|c| c.value).sum();
            },
            1 => { // Product
                value = children.iter().map(|c| c.value).fold(1, |a,b| a*b);
            },
            2 => { // Minimum
                value = children.iter().map(|c| c.value).min().expect("Couldn't find minimum");
            },
            3 => { // Maximum
                value = children.iter().map(|c| c.value).max().expect("Couldn't find maximum");
            }
            4 => { // Literal value
                let (v, l) = streamer.get_long();
                value = v;
                binary_length += l;
            },
            5 => { // Greater than
                let mut child_iter = children.iter();
                let v1 = child_iter.next().unwrap().value;
                let v2 = child_iter.next().unwrap().value;
                value = if v1 > v2 {1} else {0};
            },
            6 => { // Less than
                let mut child_iter = children.iter();
                let v1 = child_iter.next().unwrap().value;
                let v2 = child_iter.next().unwrap().value;
                value = if v1 < v2 {1} else {0};
            },
            7 => { // Equal to
                let mut child_iter = children.iter();
                let v1 = child_iter.next().unwrap().value;
                let v2 = child_iter.next().unwrap().value;
                value = if v1 == v2 {1} else {0};
            },
            _ => { panic!("Unexpected type_id {}", type_id); }
        }

        Packet {
            version,
            value,
            children,
            binary_length
        }
    }
}

struct DataStreamer {
    data: Vec<u8>,
    pos: usize
}

impl DataStreamer {
    fn new(input: String) -> Self {
        let data = Vec::from_hex(input).expect("Invalid hex string");
        DataStreamer {
            data,
            pos: 0
        }
    }

    fn get_bit(&mut self) -> u64 {
        let cur_byte = self.pos / 8;
        let cur_bit = self.pos % 8;
        let mask: u8 = 1 << (7-cur_bit);
        self.pos += 1;
        if (self.data[cur_byte] & mask) > 0 {1} else {0}
    }

    fn take(&mut self, bits: usize) -> u64 {
        let mut result = 0;
        for _ in 0..bits {
            result = result << 1;
            result += self.get_bit();
        }
        result
    }

    fn get_long(&mut self) -> (u64, usize) {
        let mut result = 0;
        let mut bit_length = 0;
        loop {
            bit_length += 5;
            let more_data_flag = self.take(1);
            result = result << 4;
            result += self.take(4);
            if more_data_flag == 0 {
                break;
            }
        }
        (result, bit_length)
    }
}