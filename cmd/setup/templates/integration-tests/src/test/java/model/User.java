package model;

public enum User {

    TPOOLE("B9492992-BC1A-5A55-363E-B2F866351E5C", "Tate", "Poole", "Piper",
            "mauris.id.sapien@Sednecmetus.edu", "0890 384 3801"), SOMEUSER();

    private String id;
    private String first_name;
    private String last_name;
    private String middle_name;
    private String email;
    private String ph_number;

    private User() {}

    private User(String id, String first_name, String last_name, String middle_name, String email, String ph_number) {
        this.id = id;
        this.first_name = first_name;
        this.last_name = last_name;
        this.middle_name = middle_name;
        this.email = email;
        this.ph_number = ph_number;
    }

    public String getId() {
        return id;
    }

    public void setId(String id) {
        this.id = id;
    }

    public String getFirst_name() {
        return first_name;
    }

    public void setFirst_name(String first_name) {
        this.first_name = first_name;
    }

    public String getLast_name() {
        return last_name;
    }

    public void setLast_name(String last_name) {
        this.last_name = last_name;
    }

    public String getMiddle_name() {
        return middle_name;
    }

    public void setMiddle_name(String middle_name) {
        this.middle_name = middle_name;
    }

    public String getEmail() {
        return email;
    }

    public void setEmail(String email) {
        this.email = email;
    }

    public String getPh_number() {
        return ph_number;
    }

    public void setPh_number(String ph_number) {
        this.ph_number = ph_number;
    }
}