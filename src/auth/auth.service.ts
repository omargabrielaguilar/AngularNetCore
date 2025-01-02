import { Injectable } from '@nestjs/common';
import { createClient } from '@supabase/supabase-js';
import * as bcrypt from 'bcrypt';

@Injectable()
export class AuthService {
  private supabase = createClient(
    process.env.SUPABASE_URL,
    process.env.SUPABASE_KEY,
  );

  async register(email: string, password: string) {
    const hashedPassword = await bcrypt.hash(password, 10);

    const { data, error } = await this.supabase.from('users').insert([
      { email, password: hashedPassword },
    ]);

    console.log(hashedPassword);

    if (error) {
      throw new Error(error.message);
    }

    return data;
  }

  async login(email: string, password: string) {
    const { data: user, error } = await this.supabase
      .from('users')
      .select('*')
      .eq('email', email)
      .single();

    if (error || !user) {
      throw new Error('Invalid credentials');
    }

    const isPasswordValid = await bcrypt.compare(password, user.password);

    if (!isPasswordValid) {
      throw new Error('Invalid credentials');
    }

    return { message: 'Login successful' };
  }
}
