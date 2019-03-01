package Utils;

import org.apache.commons.lang3.RandomStringUtils;

import java.text.ParseException;
import java.text.SimpleDateFormat;
import java.util.ArrayList;
import java.util.Date;
import java.util.List;
import java.util.Random;
import java.util.regex.Matcher;
import java.util.regex.Pattern;

public class DataGenerator {

    /** Returns a String of random numbers of desired length
     * @param length
     * @return random String of number of desired length */
    public String randomNumber(int length) {
        return RandomStringUtils.randomNumeric(length);
    }

    /** Returns a random int between min and max
     * @param min
     * @param max
     * @return random int between min and max */
    public int getRandomIntegerBetween(int min, int max) {
        Random rand = new Random();
        return rand.nextInt((max - min) + 1) + min;
    }

    /** Returns a String of random numbers and letters of desired length
     * @param length
     * @return random String of alphanumeric of desired length*/
    public String randomAlphaNumeric(int length) {
        return RandomStringUtils.randomAlphanumeric(length);
    }

    /** Returns a String of random letters of desired length
     * @param length
     * @return random String of letters of desired length*/
    public String randomAlphabetic(int length) {
        return RandomStringUtils.randomAlphabetic(length);
    }

    /** Re-formats date from originalFormat to newFormat
     *  date format patten matching found here: https://docs.oracle.com/javase/7/docs/api/java/text/SimpleDateFormat.html
     *
     * @param date
     * @param originalFormat
     * @param newFormat
     * @return date in new format */
    public static String reFormatDate(String date, String originalFormat, String newFormat) {
        String outputDate = "";
        try {
            SimpleDateFormat sdf = new SimpleDateFormat(originalFormat);
            Date tempDate = sdf.parse(date);
            SimpleDateFormat outputDateFormat = new SimpleDateFormat(newFormat);
            outputDate = outputDateFormat.format(tempDate);
        } catch (ParseException ex) {
            System.out.println("Parse Exception");
        }
        return outputDate;
    }

    /** Return a list that contains the capture groups of the regex expression
     * when compiled against the string. Index 0 is the original string. Subesequent indexes are regex-parsed groupings.
     *
     * Remember! You need "\\\\" to match a single "\" in a string. Because java and regex escapes.
     *
     * Go to https://rubular.com for amazing real-time regex
     *
     * @param string
     * @param regularExpression
     * @return list of Strings parsed by regex*/
    public List<String> getValuesFromRegEx(String string, String regularExpression) {
        List<String> returnList = new ArrayList<String>();
        Matcher m = Pattern.compile(regularExpression).matcher(string);
        if (m.find()) {
            for (int i = 0; i <= m.groupCount(); i++) {
                returnList.add(m.group(i));
            }
        }
        return returnList;
    }
}
