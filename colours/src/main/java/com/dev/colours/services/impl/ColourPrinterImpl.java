package com.dev.colours.services.impl;

import com.dev.colours.services.BluePrinter;
import com.dev.colours.services.ColourPrinter;
import com.dev.colours.services.GreenPinter;
import com.dev.colours.services.RedPrinter;

public class ColourPrinterImpl implements ColourPrinter {
    private RedPrinter redPrinter;
    private BluePrinter bluePrinter;
    private GreenPinter greenPinter;

    public ColourPrinterImpl(){
        this.redPrinter = new EnglishRedPrinter();
        this.bluePrinter = new EnglishBluePrinter();
        this.greenPinter = new EnglishGreenPrinter();
    }

    @Override
    public String print(){
        return String.join(",", redPrinter.print(), bluePrinter.print(), greenPinter.print());
    }

}
